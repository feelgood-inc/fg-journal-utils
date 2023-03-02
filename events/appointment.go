package events

const (
	AppointmentCreated      TransactionalEventName = "appointment_created"
	AppointmentUpdated      TransactionalEventName = "appointment_updated"
	AppointmentDeleted      TransactionalEventName = "appointment_deleted"
	AppointmentBooked       TransactionalEventName = "appointment_booked"
	AppointmentCanceled     TransactionalEventName = "appointment_canceled"
	AppointmentRescheduled  TransactionalEventName = "appointment_rescheduled"
	AppointmentBatchCreated TransactionalEventName = "appointment_batch_created"
)
