package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServiceName      string `env:"SERVICE_NAME" envDefault:"processing-service"`
	NatsUrl          string `env:"NATS_URL,required"`
	ProcessFileTopic string `env:"PROCESS_FILE_TOPIC" envDefault:"process_file"`
	ProcessFileQueue string `env:"PROCESS_FILE_QUEUE" envDefault:"process_file_queue"`
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
