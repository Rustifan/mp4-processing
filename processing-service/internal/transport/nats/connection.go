package nats

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
)

func Connect(cfg *config.Config, logger logger.Logger) (*nats.Conn, error) {
	options := []nats.Option{
		nats.Name(cfg.ServiceName),
		nats.ReconnectWait(5 * time.Second),
		nats.MaxReconnects(10),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.Error("Disconnected from nats: ", "error", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Info("Reconnected to nats")
		}),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			logger.Error("Nats error happened", "error", err)
		}),
	}

	return nats.Connect(cfg.NatsUrl, options...)
}
