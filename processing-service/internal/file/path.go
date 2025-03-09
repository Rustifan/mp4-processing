package file

import (
	"path"

	"github.com/rustifan/mp4-processing/processing-service/config"
)

func GetFilePath(config *config.Config, file string) string {
	baseFolderPath := config.FilesFolder
	return path.Join(baseFolderPath, file)
}

func GetProcessedFilePath(config *config.Config, file string) string {
	baseFolderPath := config.ProcessedFilesFolder
	return path.Join(baseFolderPath, file)
}
