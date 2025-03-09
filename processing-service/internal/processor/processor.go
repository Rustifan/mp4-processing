package processor

import (
	"encoding/json"

	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/file"
	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
	"github.com/rustifan/mp4-processing/processing-service/internal/parser"
	"github.com/rustifan/mp4-processing/processing-service/internal/transport/nats"
)

type Processor struct {
	log              logger.Logger
	publisher        *nats.Publisher
	config           *config.Config
	fileReaderWriter file.ReaderWriter
}

func NewProcessor(log logger.Logger, config *config.Config, fileReader file.ReaderWriter, publisher *nats.Publisher) *Processor {
	return &Processor{
		log:              log,
		config:           config,
		fileReaderWriter: fileReader,
		publisher:        publisher,
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
	filePath := object.FilePath

	processor.log.Info("Processing started for file with data", "filePath", filePath)

	filePathInContainer := file.GetFilePath(processor.config, filePath)

	fileData, err := processor.fileReaderWriter.ReadFile(filePathInContainer)
	if err != nil {
		publishProcessingFail(processor.config, processor.publisher, filePath)
		return err
	}
	initSegment, _, err := GetInitializationSegment(fileData)
	if err != nil {
		publishProcessingFail(processor.config, processor.publisher, filePath)
		return err
	}
	initSegmentPath := getInitSegmentPath(filePath)
	err = processor.fileReaderWriter.WriteFile(file.GetProcessedFilePath(processor.config, initSegmentPath), initSegment)
	processor.log.Info("Found initilization segment", "segment", initSegment)
	if err != nil {
		publishProcessingFail(processor.config, processor.publisher, filePath)
		return err
	}
	processor.log.Info("Saved initilization segment to disk")

	return publishProcessingSuccess(processor.config, processor.publisher, filePath, initSegmentPath)
}

func publishProcessingFail(config *config.Config, publisher *nats.Publisher, filePath string) error {
	dto := nats.FileUpdateTopicDto{
		FilePath: filePath,
		Status:   "Failed",
	}
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	return publisher.Publish(config.FileUpdateTopic, jsonData)
}

func publishProcessingSuccess(config *config.Config, publisher *nats.Publisher, filePath string, initSegmentPath string) error {
	dto := nats.FileUpdateTopicDto{
		FilePath:         filePath,
		Status:           "Successful",
		ProcssedFilePath: initSegmentPath,
	}
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	return publisher.Publish(config.FileUpdateTopic, jsonData)
}

func getInitSegmentPath(path string) string {
	lastDotIndex := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			lastDotIndex = i
			break
		}
	}
	if lastDotIndex == -1 {
		return path + "_init"
	}

	return path[:lastDotIndex] + "_init" + path[lastDotIndex:]
}
