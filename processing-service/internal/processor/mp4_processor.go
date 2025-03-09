package processor

import (
	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
	"github.com/rustifan/mp4-processing/processing-service/internal/parser"
	"github.com/rustifan/mp4-processing/processing-service/internal/transport/nats"
)

type Processor struct {
	log       logger.Logger
	publisher *nats.Publisher
	config    *config.Config
}

func NewProcessor(log logger.Logger, config *config.Config) *Processor {
	return &Processor{
		log:    log,
		config: config,
	}
}

func (processor *Processor) SetPublisher(publisher *nats.Publisher) *Processor {
	processor.publisher = publisher
	return processor
}

func (processor *Processor) ProcessFile(data []byte) error {
	object, err := parser.GetJSONParser[nats.ProcessFileDto]().Parse(data)
	if err != nil {
		return err
	}

	processor.log.Info("Processing started for file with data", "data", object)
	return nil
}
