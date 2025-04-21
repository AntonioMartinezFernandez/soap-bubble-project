package di

import (
	soapbubblemachineapplication "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/application"
	soapbubblemachineinfra "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/infra"
)

type SoapBubbleServices struct {
	SwitchOnSoapBubbleMachineCommandHandler  *soapbubblemachineapplication.SwitchOnSoapBubbleMachineCommandHandler
	SwitchOffSoapBubbleMachineCommandHandler *soapbubblemachineapplication.SwitchOffSoapBubbleMachineCommandHandler
}

func InitSoapBubbleServices(commonServices *CommonServices) *SoapBubbleServices {
	remoteController := soapbubblemachineinfra.NewHttpSoapBubbleMachineRemoteController()

	switchOnSoapBubbleMachineCommandHandler := soapbubblemachineapplication.NewSwitchOnSoapBubbleMachineCommandHandler(
		remoteController,
		commonServices.Repositories.SoapBubbleMachine,
	)
	switchOffSoapBubbleMachineCommandHandler := soapbubblemachineapplication.NewSwitchOffSoapBubbleMachineCommandHandler(
		remoteController,
		commonServices.Repositories.SoapBubbleMachine,
	)

	soapBubbleModuleServices := &SoapBubbleServices{
		SwitchOnSoapBubbleMachineCommandHandler:  switchOnSoapBubbleMachineCommandHandler,
		SwitchOffSoapBubbleMachineCommandHandler: switchOffSoapBubbleMachineCommandHandler,
	}

	registerSoapBubbleCommandHandlers(commonServices, soapBubbleModuleServices)

	return soapBubbleModuleServices
}

func registerSoapBubbleCommandHandlers(commonServices *CommonServices, soapBubbleModuleServices *SoapBubbleServices) {
	registerCommandOrPanic(
		commonServices.CommandBus,
		&soapbubblemachineapplication.SwitchOnSoapBubbleMachineCommand{},
		soapBubbleModuleServices.SwitchOnSoapBubbleMachineCommandHandler,
	)
	registerCommandOrPanic(
		commonServices.CommandBus,
		&soapbubblemachineapplication.SwitchOffSoapBubbleMachineCommand{},
		soapBubbleModuleServices.SwitchOffSoapBubbleMachineCommandHandler,
	)
}
