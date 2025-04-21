package soapbubblemachineapplication

import (
	"context"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
)

type SwitchOffSoapBubbleMachineCommandHandler struct {
	soapBubbleMachineRemoteController soapbubblemachinedomain.SoapBubbleMachineRemoteController
	soapBubbleMachineRepository       soapbubblemachinedomain.SoapBubbleMachineRepository
}

func NewSwitchOffSoapBubbleMachineCommandHandler(
	soapbubblemachineremotecontroller soapbubblemachinedomain.SoapBubbleMachineRemoteController,
	soapbubblemachinerepository soapbubblemachinedomain.SoapBubbleMachineRepository,
) *SwitchOffSoapBubbleMachineCommandHandler {
	return &SwitchOffSoapBubbleMachineCommandHandler{
		soapBubbleMachineRemoteController: soapbubblemachineremotecontroller,
		soapBubbleMachineRepository:       soapbubblemachinerepository,
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

	sbm, err := h.soapBubbleMachineRepository.FindByIdentifier(
		ctx,
		switchOffSoapBubbleMachineCommand.Namespace,
		switchOffSoapBubbleMachineCommand.SoapBubbleMachineID,
	)
	if err != nil {
		return err
	}

	return sbm.SwitchOFF(ctx, h.soapBubbleMachineRemoteController)
}
