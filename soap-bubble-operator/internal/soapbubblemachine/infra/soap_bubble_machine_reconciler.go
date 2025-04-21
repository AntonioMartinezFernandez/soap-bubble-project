/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package soapbubblemachineinfra

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	soapbubbleoperatorv1alpha1 "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/api/v1alpha1"
	soapbubblemachineapplication "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/application"
	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus/command"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/logger"
)

const (
	soapBubbleMachineControllerFinalizer = "antonio.mf/soap-bubble-machine-operator"
)

// SoapBubbleMachineReconciler reconciles a SoapBubbleMachine object
type SoapBubbleMachineReconciler struct {
	k8sClient  client.Client
	scheme     *runtime.Scheme
	commandBus command.Bus
	logger     logger.Logger
}

func NewSoapBubbleMachineReconciler(
	k8sClient client.Client,
	scheme *runtime.Scheme,
	commandBus command.Bus,
	logger logger.Logger,
) *SoapBubbleMachineReconciler {
	return &SoapBubbleMachineReconciler{
		k8sClient:  k8sClient,
		scheme:     scheme,
		commandBus: commandBus,
		logger:     logger,
	}
}

// +kubebuilder:rbac:groups=soap-bubble-operator.soap-bubble-operator.local,resources=soapbubblemachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=soap-bubble-operator.soap-bubble-operator.local,resources=soapbubblemachines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=soap-bubble-operator.soap-bubble-operator.local,resources=soapbubblemachines/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SoapBubbleMachine object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *SoapBubbleMachineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reconcileID := string(controller.ReconcileIDFromContext(ctx))

	var soapBubbleMachine soapbubbleoperatorv1alpha1.SoapBubbleMachine
	if err := r.k8sClient.Get(ctx, req.NamespacedName, &soapBubbleMachine); err != nil {
		r.logger.Info(
			ctx,
			"⚠️ unable to fetch soap bubble machine",
			slog.String("name", req.Name),
			slog.String("namespace", req.Namespace),
			slog.String("reconcileID", reconcileID),
		)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.logger.Info(
		ctx,
		"⚙️ reconciling soap bubble machine",
		slog.String("name", req.Name),
		slog.String("namespace", req.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	if !soapBubbleMachine.GetDeletionTimestamp().IsZero() {
		//! Process resource deletion
		if err := r.deleteSoapBubbleMachine(ctx, reconcileID, soapBubbleMachine); err != nil {
			r.logger.Error(
				ctx,
				"❌ error during deletion",
				slog.String("name", soapBubbleMachine.Name),
				slog.String("namespace", req.Namespace),
				slog.String("reconcileID", reconcileID),
				slog.String("error", err.Error()),
			)

			return ctrl.Result{RequeueAfter: 10 * time.Second}, err
		}

		return ctrl.Result{}, nil
	}

	if controllerutil.AddFinalizer(&soapBubbleMachine, soapBubbleMachineControllerFinalizer) {
		if err := r.k8sClient.Update(ctx, &soapBubbleMachine); err != nil {
			r.logger.Error(
				ctx,
				"❌ unable to add finalizer",
				slog.String("name", soapBubbleMachine.Name),
				slog.String("namespace", req.Namespace),
				slog.String("reconcileID", reconcileID),
				slog.String("error", err.Error()),
			)
			return ctrl.Result{}, fmt.Errorf("adding finalizer: %w", err)
		}

		return ctrl.Result{Requeue: true}, nil
	}

	//! Process resource upsert
	err := r.upsertSoapBubbleMachine(ctx, reconcileID, soapBubbleMachine)
	if err != nil {
		r.logger.Error(
			ctx,
			"❌ error during upsert",
			slog.String("name", soapBubbleMachine.Name),
			slog.String("namespace", req.Namespace),
			slog.String("reconcileID", reconcileID),
			slog.String("error", err.Error()),
		)
		return ctrl.Result{RequeueAfter: 3 * time.Second}, err
	}

	r.logger.Info(
		ctx,
		"🎉 soap bubble machine reconciled",
		slog.String("name", soapBubbleMachine.Name),
		slog.String("namespace", req.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SoapBubbleMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	predicate := predicate.GenerationChangedPredicate{} // Disable reconciliation on status changes
	return ctrl.NewControllerManagedBy(mgr).
		For(&soapbubbleoperatorv1alpha1.SoapBubbleMachine{}).
		WithEventFilter(predicate).
		Complete(r)
}

func (r *SoapBubbleMachineReconciler) deleteSoapBubbleMachine(ctx context.Context, reconcileID string, soapBubbleMachine soapbubbleoperatorv1alpha1.SoapBubbleMachine) error {
	r.logger.Info(
		ctx,
		"⏳ deleting soap bubble machine",
		slog.String("name", soapBubbleMachine.Name),
		slog.String("namespace", soapBubbleMachine.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	switchOffSoapBubbleMachineCommand := soapbubblemachineapplication.NewSwitchOffSoapBubbleMachineCommand(
		soapbubblemachinedomain.NewSoapBubbleMachineID(soapBubbleMachine.Namespace, soapBubbleMachine.Name).String(),
		soapBubbleMachine.Name,
		soapBubbleMachine.Spec.IP,
	)

	if err := r.commandBus.Exec(ctx, switchOffSoapBubbleMachineCommand); err != nil {
		return fmt.Errorf("failed switching off soap bubble machine: %w", err)
	}

	if controllerutil.RemoveFinalizer(&soapBubbleMachine, soapBubbleMachineControllerFinalizer) {
		if err := r.k8sClient.Update(ctx, &soapBubbleMachine); err != nil {
			r.logger.Error(
				ctx,
				"❌ failed to remove finalizer",
				slog.String("name", soapBubbleMachine.Name),
				slog.String("error", err.Error()),
			)
			return err
		}
	}

	r.logger.Info(
		ctx,
		"✅ soap bubble machine deleted",
		slog.String("name", soapBubbleMachine.Name),
		slog.String("namespace", soapBubbleMachine.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	return nil
}

func (r *SoapBubbleMachineReconciler) upsertSoapBubbleMachine(
	ctx context.Context,
	reconcileID string,
	soapBubbleMachine soapbubbleoperatorv1alpha1.SoapBubbleMachine,
) error {
	r.logger.Info(
		ctx,
		"⏳ upserting soap bubble machine",
		slog.String("name", soapBubbleMachine.Name),
		slog.String("namespace", soapBubbleMachine.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	// Switch on the soap bubble machine
	if soapBubbleMachine.Spec.MakeBubbles && !soapBubbleMachine.Status.MakingBubbles {
		switchOnSoapBubbleMachineCommand := soapbubblemachineapplication.NewSwitchOnSoapBubbleMachineCommand(
			soapbubblemachinedomain.NewSoapBubbleMachineID(soapBubbleMachine.Namespace, soapBubbleMachine.Name).String(),
			soapBubbleMachine.Name,
			soapBubbleMachine.Spec.IP,
			soapBubbleMachine.Spec.Speed,
		)

		if err := r.commandBus.Exec(ctx, switchOnSoapBubbleMachineCommand); err != nil {
			return fmt.Errorf("failed switching on soap bubble machine: %w", err)
		}

		soapBubbleMachineCopy := soapBubbleMachine.DeepCopy() // avoid mutating the cached object
		soapBubbleMachineCopy.Status.MakingBubbles = true

		if err := r.k8sClient.Status().Update(ctx, soapBubbleMachineCopy); err != nil {
			return err
		}
	}

	// Switch off the soap bubble machine
	if !soapBubbleMachine.Spec.MakeBubbles && soapBubbleMachine.Status.MakingBubbles {
		switchOffSoapBubbleMachineCommand := soapbubblemachineapplication.NewSwitchOffSoapBubbleMachineCommand(
			soapbubblemachinedomain.NewSoapBubbleMachineID(soapBubbleMachine.Namespace, soapBubbleMachine.Name).String(),
			soapBubbleMachine.Name,
			soapBubbleMachine.Spec.IP,
		)

		if err := r.commandBus.Exec(ctx, switchOffSoapBubbleMachineCommand); err != nil {
			return fmt.Errorf("failed switching off soap bubble machine: %w", err)
		}

		soapBubbleMachineCopy := soapBubbleMachine.DeepCopy() // avoid mutating the cached object
		soapBubbleMachineCopy.Status.MakingBubbles = false

		if err := r.k8sClient.Status().Update(ctx, soapBubbleMachineCopy); err != nil {
			return err
		}
	}

	r.logger.Info(
		ctx,
		"✅ soap bubble machine upsert finished",
		slog.String("name", soapBubbleMachine.Name),
		slog.String("namespace", soapBubbleMachine.Namespace),
		slog.String("reconcileID", reconcileID),
	)

	return nil
}
