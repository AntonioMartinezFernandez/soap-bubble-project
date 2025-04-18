package soapbubblemachineapplication

import (
	"context"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
)

type SwitchOffSoapBubbleMachineCommandHandler struct {
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController
}

func NewSwitchOffSoapBubbleMachineCommandHandler(
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController,
) *SwitchOffSoapBubbleMachineCommandHandler {
	return &SwitchOffSoapBubbleMachineCommandHandler{
		soapbubblemachineremotecontroller: soapbubblemachineremotecontroller,
	}
}

func (h *SwitchOffSoapBubbleMachineCommandHandler) Handle(
	ctx context.Context,
	command bus.Dto,
) error {
	switchOffSoapBubbleMachineCommand, ok := command.(*SwitchOffSoapBubbleMachineCommand)
	if !ok {
		return bus.NewInvalidDto("expected SwitchOffSoapBubbleMachineCommand")
	}

	soapBubbleMachine := soapbubblemachinedomain.NewSoapBubbleMachine(
		switchOffSoapBubbleMachineCommand.SoapBubbleMachineID,
		switchOffSoapBubbleMachineCommand.SoapBubbleMachineName,
		switchOffSoapBubbleMachineCommand.StartURL,
		switchOffSoapBubbleMachineCommand.StopURL,
		switchOffSoapBubbleMachineCommand.MakingBubbles,
	)

	return h.soapbubblemachineremotecontroller.SwitchOff(ctx, *soapBubbleMachine)
}
