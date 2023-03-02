package events

import "time"

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
	OriginalResource  string                   `json:"original_resource" validate:"required_if=ActionTaken update,delete"`
	ResultingResource string                   `json:"resulting_resource" validate:"required_if=ActionTaken update,delete,create"`
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