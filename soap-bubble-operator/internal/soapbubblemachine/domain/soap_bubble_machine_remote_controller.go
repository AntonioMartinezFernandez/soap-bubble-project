package soapbubblemachinedomain

import "context"

type SoapBubbleMachineRemoteController interface {
	SwitchOn(ctx context.Context, soapBubbleMachine SoapBubbleMachine) error
	SwitchOff(ctx context.Context, soapBubbleMachine SoapBubbleMachine) error
}
