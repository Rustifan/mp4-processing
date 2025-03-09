package nats

import (
	"context"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
)

type MessageHandler func([]byte) error

type Subscriber struct {
	connection    *nats.Conn
	log           logger.Logger
	subscriptions []*nats.Subscription
	mutex         sync.Mutex
}

func NewSubscriber(connection *nats.Conn, logger logger.Logger) *Subscriber {
	return &Subscriber{
		connection:    connection,
		log:           logger,
		subscriptions: make([]*nats.Subscription, 0),
	}
}

func (subscriber *Subscriber) QueueSubscribe(subject string, queue string, handler MessageHandler) error {
	subscriber.mutex.Lock()
	defer subscriber.mutex.Unlock()

	sub, err := subscriber.connection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		if err := handler(msg.Data); err != nil {
			subscriber.log.Error("Error processing data", "error", err)
		}
	})

	if err != nil {
		return err
	}
	subscriber.subscriptions = append(subscriber.subscriptions, sub)
	subscriber.log.Info("Subscribed to subject", "subject", subject)

	return nil
}

func (sunscriber *Subscriber) Shutdown(ctx context.Context) error {
	sunscriber.mutex.Lock()
	defer sunscriber.mutex.Unlock()

	for _, sub := range sunscriber.subscriptions {
		if err := sub.Unsubscribe(); err != nil {
			sunscriber.log.Error("Error unsubscribing", "error", err)
		}
	}

	return nil
}
