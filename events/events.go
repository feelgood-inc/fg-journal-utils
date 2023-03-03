package events

import (
	"github.com/invopop/validation"
	"time"
)

type TransactionalEventName string
type TransactionalEventAction string

const (
	TransactionalEventActionCreate TransactionalEventAction = "create"
	TransactionalEventActionUpdate TransactionalEventAction = "update"
	TransactionalEventActionDelete TransactionalEventAction = "delete"
)

type TransactionalEventMetadata struct {
	TransactionID string    `json:"transaction_id"`
	SentAt        time.Time `json:"sent_at"`
	SentBy        string    `json:"sent_by"`
}

type TransactionalEventResource struct {
	OriginalResource  interface{}              `json:"original_resource"`
	ResultingResource interface{}              `json:"resulting_resource"`
	ActionTaken       TransactionalEventAction `json:"action_taken"`
}

type TransactionEvent struct {
	Event    TransactionalEventName     `json:"event"`
	Metadata TransactionalEventMetadata `json:"metadata"`
	Resource TransactionalEventResource `json:"resource"`
}

type TransactionalEventPayload struct {
	Event    TransactionalEventName     `json:"event" validate:"required"`
	Resource TransactionalEventResource `json:"resource" validate:"required"`
	SentBy   string                     `json:"sent_by" validate:"required"`
}

type TransactionalEventClientConfig struct {
	Host  string
	Debug bool
}

func (c TransactionEvent) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Event, validation.Required),
		validation.Field(&c.Metadata),
		validation.Field(&c.Resource, validation.Required),
	)
}

func (c TransactionalEventMetadata) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.TransactionID, validation.Required, validation.Length(21, 21)),
		validation.Field(&c.SentAt, validation.Required),
		validation.Field(&c.SentBy, validation.Required),
	)
}

func (c TransactionalEventPayload) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Event, validation.Required),
		validation.Field(&c.Resource),
		validation.Field(&c.SentBy, validation.Required),
	)
}

func (c TransactionalEventResource) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.OriginalResource,
			validation.When(c.ActionTaken == TransactionalEventActionUpdate ||
				c.ActionTaken == TransactionalEventActionDelete, validation.Required)),
		validation.Field(&c.ResultingResource,
			validation.When(c.ActionTaken == TransactionalEventActionUpdate ||
				c.ActionTaken == TransactionalEventActionDelete ||
				c.ActionTaken == TransactionalEventActionCreate, validation.Required)),
		validation.Field(&c.ActionTaken, validation.Required, validation.In(TransactionalEventActionCreate, TransactionalEventActionUpdate, TransactionalEventActionDelete)),
	)
}
