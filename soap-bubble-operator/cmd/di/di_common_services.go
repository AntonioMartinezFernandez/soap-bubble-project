package di

import (
	"context"

	soapbubblemachinedomain "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/domain"
	soapbubblemachineinfra "github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/internal/soapbubblemachine/infra"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus/command"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/logger"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Repositories struct {
	SoapBubbleMachine soapbubblemachinedomain.SoapBubbleMachineRepository
}

type CommonServices struct {
	RootCtx    context.Context
	CommandBus command.Bus
	Logger     logger.Logger

	Repositories *Repositories
}

func StartCommonServices(k8sClient client.Client) *CommonServices {
	ctx := context.Background()
	l := logger.NewLogger("debug")
	repositories := StartRepositories(k8sClient)

	return &CommonServices{
		RootCtx:    ctx,
		CommandBus: command.InitCommandBus(l),
		Logger:     l,

		Repositories: repositories,
	}
}

func StartRepositories(k8sClient client.Client) *Repositories {
	soapBubbleMachineRepository := soapbubblemachineinfra.NewK8sSoapBubbleMachineRepository(k8sClient)

	return &Repositories{
		SoapBubbleMachine: soapBubbleMachineRepository,
	}
}

func registerCommandOrPanic(
	commandBus command.Bus,
	cmd bus.Dto,
	handler command.CommandHandler,
) {
	if err := commandBus.RegisterCommand(cmd, handler); err != nil {
		panic(err)
	}
}
