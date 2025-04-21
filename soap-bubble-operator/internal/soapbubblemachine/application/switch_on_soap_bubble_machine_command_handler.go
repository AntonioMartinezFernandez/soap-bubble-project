package soapbubblemachineapplication

import (
	"context"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
)

type SwitchOnSoapBubbleMachineCommandHandler struct {
	soapBubbleMachineRemoteController soapbubblemachinedomain.SoapBubbleMachineRemoteController
	soapBubbleMachineRepository       soapbubblemachinedomain.SoapBubbleMachineRepository
}

func NewSwitchOnSoapBubbleMachineCommandHandler(
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController,
	soapbubblemachinerepository soapbubblemachinedomain.SoapBubbleMachineRepository,
) *SwitchOnSoapBubbleMachineCommandHandler {
	return &SwitchOnSoapBubbleMachineCommandHandler{
		soapBubbleMachineRemoteController: soapbubblemachineremotecontroller,
		soapBubbleMachineRepository:       soapbubblemachinerepository,
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

	sbm, err := h.soapBubbleMachineRepository.FindByIdentifier(
		ctx,
		switchOnSoapBubbleMachineCommand.Namespace,
		switchOnSoapBubbleMachineCommand.SoapBubbleMachineID,
	)
	if err != nil {
		return err
	}

	sbm.SetSpeed(switchOnSoapBubbleMachineCommand.Speed)

	return sbm.SwitchON(ctx, h.soapBubbleMachineRemoteController)
}
