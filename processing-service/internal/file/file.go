package file

import (
	"os"

	"github.com/rustifan/mp4-processing/processing-service/internal/logger"
)

type ReaderWriter interface {
	ReadFile(filePath string) ([]byte, error)
	WriteFile(filePath string, data []byte) error
}

type FileReaderWriter struct {
	log logger.Logger
}

func NewFileReader(logger logger.Logger) *FileReaderWriter {
	return &FileReaderWriter{log: logger}
}

func (reader *FileReaderWriter) ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (writer *FileReaderWriter) WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
