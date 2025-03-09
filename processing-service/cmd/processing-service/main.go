package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
	"github.com/rustifan/mp4-processing/processing-service/internal/processor"
	"github.com/rustifan/mp4-processing/processing-service/internal/transport/nats"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error while getting config", err)
		os.Exit(1)
	}

	logger, err := logger.NewLogger(cfg.ServiceName)
	if err != nil {
		log.Fatal("Failed to create logger", "error", err)
		os.Exit(1)
	}
	logger.Info("Configuration loaded")

	natsConn, err := nats.Connect(cfg, logger)
	if err != nil {
		log.Fatal("Failed to connect to NATS", "error", err)
		os.Exit(1)
	}
	defer natsConn.Close()

	proc := processor.NewProcessor(logger, cfg)

	subscriber := nats.NewSubscriber(natsConn, logger)
	subscribeError := subscriber.QueueSubscribe(cfg.ProcessFileTopic, cfg.ProcessFileQueue, proc.ProcessFile)
	if subscribeError != nil {
		logger.Error("Error happened while subscribing to ", "topic", cfg.ProcessFileTopic)
		os.Exit(1)
	}
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	sig := <-shutdown
	logger.Info("Received shutdown signal", "signal", sig)
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := subscriber.Shutdown(shutdownCtx); err != nil {
		logger.Error("Error during subscriber shutdown", "error", err)
	}

	logger.Info("Worker service stopped")
}
