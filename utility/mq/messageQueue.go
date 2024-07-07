package mq

func Publish(topic_ string, message_ any) error {
	return inst.Publish(topic_, message_)
}

func Subscribe(topic_ string, callback_ Callback) {
	inst.Subscribe(topic_, callback_)
}
