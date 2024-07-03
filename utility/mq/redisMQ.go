package mq

import (
	"context"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/go-redis/redis/v8"
	"sync"
)

type CallbackList = []Callback

type redisMQ struct {
	client      *redis.Client
	subscribers map[string]CallbackList

	mtx sync.RWMutex
}

func NewRedisMQ() IMQ {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolFIFO: true,
	})

	return &redisMQ{
		client:      cli,
		subscribers: make(map[string]CallbackList),
	}
}

func (r *redisMQ) Publish(topic_ string, message_ any) error {
	ctx := context.Background()
	err := r.client.Publish(ctx, topic_, message_).Err()

	if err != nil {
		log.Warnf("Failed to publish message to topic %s: %v", topic_, err)
		return err
	}
	return nil
}

func (r *redisMQ) handleTopic(topic_ string) {
	subscriber := r.client.Subscribe(context.Background(), topic_)
	ch := subscriber.Channel()
	for msg := range ch {
		r.processMessages(topic_, msg.Payload)
	}
}
func (r *redisMQ) processMessages(topic_, msg string) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	for _, cb := range r.subscribers[topic_] {
		go cb(msg)
	}
}

func (r *redisMQ) Subscribe(topic_ string, callback_ Callback) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, exists := r.subscribers[topic_]; !exists {
		r.subscribers[topic_] = CallbackList{}

		go r.handleTopic(topic_)
	}
	r.subscribers[topic_] = append(r.subscribers[topic_], callback_)

}

func (r *redisMQ) Consume(topic_ string) (any, error) {
	return r.client.Get(context.Background(), topic_).Result()
}

func (r *redisMQ) Close() error {
	statu := r.client.FlushAll(context.Background())
	log.Infof("Closing connection to redis: %v", statu)
	log.Debugf("Closing connection to redis: %s", statu.String())
	return statu.Err()
}
