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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	soapbubbleoperatorv1alpha1 "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/api/v1alpha1"
)

const (
	soapBubbleMachineControllerFinalizer = "antonio.mf/soap-bubble-machine-operator"
)

// SoapBubbleMachineReconciler reconciles a SoapBubbleMachine object
type SoapBubbleMachineReconciler struct {
	client.Client
	Scheme *runtime.Scheme
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
	log := logf.FromContext(ctx)
	reconcileID := string(controller.ReconcileIDFromContext(ctx))

	var soapBubbleMachine soapbubbleoperatorv1alpha1.SoapBubbleMachine
	if err := r.Client.Get(ctx, req.NamespacedName, &soapBubbleMachine); err != nil {
		log.Info("unable to fetch soapBubbleMachine", "name", req.Name)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("reconciling soapBubbleMachinee", "name", soapBubbleMachine.Name)

	if !soapBubbleMachine.GetDeletionTimestamp().IsZero() {

		//! Process resource deletion
		err := dispatchResourceDeletion(reconcileID)
		if err != nil {
			log.Error(err, "error during deletion", "name", soapBubbleMachine.Name)
			return ctrl.Result{}, err
		}

		if controllerutil.RemoveFinalizer(&soapBubbleMachine, soapBubbleMachineControllerFinalizer) {
			if err := r.Update(ctx, &soapBubbleMachine); err != nil {
				log.Error(err, "failed to remove finalizer from soapBubbleMachine", "name", soapBubbleMachine.Name)
				return ctrl.Result{}, err
			}
		}

		log.Info("soapBubbleMachine finalizer removed", "name", soapBubbleMachine.Name)
		return ctrl.Result{}, nil
	}

	if controllerutil.AddFinalizer(&soapBubbleMachine, soapBubbleMachineControllerFinalizer) {
		if err := r.Update(ctx, &soapBubbleMachine); err != nil {
			log.Error(err, "unable to add finalizer")
			return ctrl.Result{}, fmt.Errorf("adding finalizer: %w", err)
		}

		return ctrl.Result{Requeue: true}, nil
	}

	//! Process resource upsert
	err := dispatchResourceUpsert(reconcileID)
	if err != nil {
		log.Error(err, "error during upsert", "name", soapBubbleMachine.Name)
		return ctrl.Result{}, err
	}

	log.Info("soapBubbleMachine reconciled", "Namespace", req.Namespace, "Name", req.Name)

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

func dispatchResourceDeletion(uniqueID string) error {
	fmt.Printf("⏳ EXECUTING DELETE 'SOAP BUBBLE MACHINE' PROCESS %s\n", uniqueID)

	//! TODO: Switch off the machine and delete the resource

	fmt.Printf("✅ DELETE 'SOAP BUBBLE MACHINE' PROCESS %s FINISHED\n", uniqueID)
	return nil
}

func dispatchResourceUpsert(uniqueID string) error {
	fmt.Printf("⏳ EXECUTING CREATE/UPDATE 'SOAP BUBBLE MACHINE' PROCESS %s\n", uniqueID)

	//! TODO: Get the desired state of the machine, switch on/off the machine and create/update the resource

	fmt.Printf("✅ CREATE/UPDATE 'SOAP BUBBLE MACHINE' PROCESS %s FINISHED\n", uniqueID)
	return nil
}
