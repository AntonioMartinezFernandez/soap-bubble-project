package command

import (
	"context"

	"github.com/AntonioMartinezFernandez/soap-bubble-project/soap-bubble-operator/pkg/bus"
)

type CommandHandler interface {
	Handle(ctx context.Context, command bus.Dto) error
}
