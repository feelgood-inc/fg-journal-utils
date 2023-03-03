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
	TransactionID string    `json:"transaction_id" bson:"transactionID"`
	SentAt        time.Time `json:"sent_at" bson:"sentAt"`
	SentBy        string    `json:"sent_by" bson:"sentBy"`
	ExecutedByUID string    `json:"executed_by_uid" bson:"executedByUID"`
}

type TransactionalEventResource struct {
	OriginalResource  interface{}              `json:"original_resource" bson:"originalResource"`
	ResultingResource interface{}              `json:"resulting_resource" bson:"resultingResource"`
	ActionTaken       TransactionalEventAction `json:"action_taken" bson:"actionTaken"`
	Name              string                   `json:"name" bson:"name"`
}

type TransactionEvent struct {
	Event    TransactionalEventName     `json:"event" bson:"event"`
	Metadata TransactionalEventMetadata `json:"metadata" bson:"metadata"`
	Resource TransactionalEventResource `json:"resource" bson:"resource"`
}

type TransactionalEventPayload struct {
	Event         TransactionalEventName     `json:"event" validate:"required" bson:"event"`
	Resource      TransactionalEventResource `json:"resource" validate:"required" bson:"resource"`
	SentBy        string                     `json:"sent_by" validate:"required" bson:"sentBy"`
	ExecutedByUID string                     `json:"executed_by_uid" bson:"executedByUID"`
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
		validation.Field(&c.Name, validation.Required),
	)
}
