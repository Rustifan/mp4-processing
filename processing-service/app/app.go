package app

import (
	"context"

	natsPackage "github.com/nats-io/nats.go"
	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/file"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
	"github.com/rustifan/mp4-processing/processing-service/internal/processor"
	"github.com/rustifan/mp4-processing/processing-service/internal/transport/nats"
)

type App struct {
	cfg        *config.Config
	logger     logger.Logger
	natsConn   *natsPackage.Conn
	subscriber *nats.Subscriber
	publisher  *nats.Publisher
	processor  *processor.Processor
}

func NewApp() (*App, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	logger, err := logger.NewLogger(cfg.ServiceName)
	if err != nil {
		return nil, err
	}
	logger.Info("Configuration loaded")

	natsConn, err := nats.Connect(cfg, logger)
	if err != nil {
		return nil, err
	}

	publisher := nats.NewPublisher(natsConn, logger)
	fileReader := file.NewFileReader(logger)
	proc := processor.NewProcessor(logger, cfg, fileReader, publisher)
	subscriber := nats.NewSubscriber(natsConn, logger)

	return &App{
		cfg:        cfg,
		logger:     logger,
		natsConn:   natsConn,
		subscriber: subscriber,
		publisher:  publisher,
		processor:  proc,
	}, nil
}

func (a *App) Start() error {
	subscribeError := a.subscriber.QueueSubscribe(
		a.cfg.ProcessFileTopic,
		a.cfg.ProcessFileQueue,
		a.processor.ProcessFile,
	)
	if subscribeError != nil {
		a.logger.Error("Error happened while subscribing to ", "topic", a.cfg.ProcessFileTopic)
		return subscribeError
	}

	a.logger.Info("Application started successfully")
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Shutting down application...")

	if err := a.subscriber.Shutdown(ctx); err != nil {
		a.logger.Error("Error during subscriber shutdown", "error", err)
		return err
	}

	a.natsConn.Close()
	a.logger.Info("Application stopped successfully")
	return nil
}
