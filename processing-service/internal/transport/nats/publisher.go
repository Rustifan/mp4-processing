package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
)

type Publisher struct {
	connection *nats.Conn
	log        logger.Logger
}

func NewPublisher(connection *nats.Conn, logger logger.Logger) *Publisher {
	return &Publisher{
		connection: connection,
		log:        logger,
	}
}

func (publisher *Publisher) Publish(subject string, data []byte) error {
	if err := publisher.connection.Publish(subject, data); err != nil {
		publisher.log.Error("Failed to publish on subject", "subject", subject, "error", err)
		return err
	}

	publisher.log.Info("Sucesfully published on subject", "subject", subject)
	return nil
}
