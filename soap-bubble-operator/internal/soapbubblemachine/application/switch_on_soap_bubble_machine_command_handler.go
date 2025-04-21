package soapbubblemachineapplication

import (
	"context"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
)

type SwitchOnSoapBubbleMachineCommandHandler struct {
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController
}

func NewSwitchOnSoapBubbleMachineCommandHandler(
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController,
) *SwitchOnSoapBubbleMachineCommandHandler {
	return &SwitchOnSoapBubbleMachineCommandHandler{
		soapbubblemachineremotecontroller: soapbubblemachineremotecontroller,
	}
}

func (h *SwitchOnSoapBubbleMachineCommandHandler) Handle(
	ctx context.Context,
	command bus.Dto,
) error {
	switchOnSoapBubbleMachineCommand, ok := command.(*SwitchOnSoapBubbleMachineCommand)
	if !ok {
		return bus.NewInvalidDto("expected SwitchOnSoapBubbleMachineCommand")
	}

	soapBubbleMachine := soapbubblemachinedomain.NewSoapBubbleMachine(
		switchOnSoapBubbleMachineCommand.SoapBubbleMachineID,
		switchOnSoapBubbleMachineCommand.SoapBubbleMachineName,
		switchOnSoapBubbleMachineCommand.SoapBubbleMachineIP,
		true,
		switchOnSoapBubbleMachineCommand.Speed,
	)

	return soapBubbleMachine.SwitchON(ctx, h.soapbubblemachineremotecontroller)
}
