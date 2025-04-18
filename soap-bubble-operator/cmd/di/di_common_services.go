package di

import (
	"context"

	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus/command"
	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/logger"
)

type CommonServices struct {
	RootCtx    context.Context
	CommandBus command.Bus
	Logger     logger.Logger
}

func StartCommonServices() *CommonServices {
	ctx := context.Background()
	l := logger.NewLogger("debug")

	return &CommonServices{
		RootCtx:    ctx,
		CommandBus: command.InitCommandBus(l),
		Logger:     l,
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
