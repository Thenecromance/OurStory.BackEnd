package mq

var inst IMQ

func init() {
	inst = NewRedisMQ()
}

func Instance() IMQ {
	return inst
}
