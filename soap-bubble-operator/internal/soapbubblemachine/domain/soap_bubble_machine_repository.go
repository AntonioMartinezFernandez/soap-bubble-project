package soapbubblemachinedomain

import "context"

type SoapBubbleMachineRepository interface {
	FindByIdentifier(ctx context.Context, namespace, identifier string) (*SoapBubbleMachine, error)
}
