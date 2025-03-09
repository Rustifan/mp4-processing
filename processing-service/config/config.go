package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServiceName          string `env:"SERVICE_NAME" envDefault:"processing-service"`
	NatsUrl              string `env:"NATS_URL,required"`
	ProcessFileTopic     string `env:"PROCESS_FILE_TOPIC" envDefault:"process_file"`
	ProcessFileQueue     string `env:"PROCESS_FILE_QUEUE" envDefault:"process_file_queue"`
	FilesFolder          string `env:"FILES_FOLDER" envDefault:"/files"`
	ProcessedFilesFolder string `env:"PROCESSED_FILES_FOLDER" envDefault:"/processed-files"`
	FileUpdateTopic      string `env:"FILE_UPDATE_TOPIC" envDefault:"update_file"`
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
