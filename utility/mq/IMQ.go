package mq

type Callback func(message string)

type IMQ interface {
	// Publish a message to the queue
	Publish(topic_ string, message_ any) error

	Subscribe(topic_ string, callback_ Callback)
	// Consume a message from the queue
	Consume(topic_ string) (any, error)
	// Close the connection to the queue
	Close() error
}
