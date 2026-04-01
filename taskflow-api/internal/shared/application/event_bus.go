package application

import (
	"context"

	"github.com/Floxtouille/taskflow-popomagico/taskflow-api/internal/shared/domain"
)

// EventBus est le port de publication des événements métier.
// L'implémentation (NATS, mémoire, etc.) vit dans infrastructure/.
type EventBus interface {
	Publish(ctx context.Context, event domain.DomainEvent) error
}
