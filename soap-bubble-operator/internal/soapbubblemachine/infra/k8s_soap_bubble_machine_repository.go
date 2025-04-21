package soapbubblemachineinfra

import (
	"context"

	soapbubbleoperatorv1alpha1 "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/api/v1alpha1"
	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ soapbubblemachinedomain.SoapBubbleMachineRepository = (*K8sSoapBubbleMachineRepository)(nil)

type K8sSoapBubbleMachineRepository struct {
	k8sClient client.Client
}

func NewK8sSoapBubbleMachineRepository(k8sClient client.Client) *K8sSoapBubbleMachineRepository {
	return &K8sSoapBubbleMachineRepository{
		k8sClient: k8sClient,
	}
}

func (r *K8sSoapBubbleMachineRepository) FindByIdentifier(ctx context.Context, namespace, identifier string) (*soapbubblemachinedomain.SoapBubbleMachine, error) {
	var soapBubbleMachines soapbubbleoperatorv1alpha1.SoapBubbleMachineList

	reqLabel, err := labels.NewRequirement("machine/identifier", selection.Equals, []string{identifier})
	if err != nil {
		return nil, err
	}

	listOpts := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: labels.NewSelector().Add(*reqLabel),
	}

	if err := r.k8sClient.List(ctx, &soapBubbleMachines, listOpts); err != nil {
		return nil, err
	}

	var soapBubbleMachine *soapbubblemachinedomain.SoapBubbleMachine
	for _, sbm := range soapBubbleMachines.Items {
		soapBubbleMachine = soapbubblemachinedomain.NewSoapBubbleMachine(
			identifier,
			sbm.Spec.MachineName,
			sbm.Spec.IP,
			sbm.Status.MakingBubbles,
			sbm.Spec.Speed,
		)
	}

	return soapBubbleMachine, nil
}
