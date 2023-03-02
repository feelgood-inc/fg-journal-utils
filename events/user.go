package events

const (
	UserCreated                TransactionalEventName = "user_created"
	UserUpdated                TransactionalEventName = "user_updated"
	UserDeleted                TransactionalEventName = "user_deleted"
	UserActivated              TransactionalEventName = "user_activated"
	UserDeactivated            TransactionalEventName = "user_deactivated"
	UserPasswordChanged        TransactionalEventName = "user_password_changed"
	UserPasswordReset          TransactionalEventName = "user_password_reset"
	UserPasswordResetRequested TransactionalEventName = "user_password_reset_requested"
)
