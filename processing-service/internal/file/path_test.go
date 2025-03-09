package file_test

import (
	"testing"

	"github.com/rustifan/mp4-processing/processing-service/config"
	"github.com/rustifan/mp4-processing/processing-service/internal/file"
)

func TestGetFilePath(t *testing.T) {
	tests := []struct {
		name           string
		config         *config.Config
		file           string
		expectedResult string
	}{
		{
			name: "Simple file path",
			config: &config.Config{
				FilesFolder: "/test/folder",
			},
			file:           "file.mp4",
			expectedResult: "/test/folder/file.mp4",
		},
		{
			name: "File path with subdirectory",
			config: &config.Config{
				FilesFolder: "/test/folder",
			},
			file:           "subdir/file.mp4",
			expectedResult: "/test/folder/subdir/file.mp4",
		},
		{
			name: "Path with trailing slash in config",
			config: &config.Config{
				FilesFolder: "/test/folder/",
			},
			file:           "file.mp4",
			expectedResult: "/test/folder/file.mp4",
		},
		{
			name: "Path with relative file notation",
			config: &config.Config{
				FilesFolder: "/test/folder",
			},
			file:           "./file.mp4",
			expectedResult: "/test/folder/file.mp4",
		},
		{
			name: "Empty file path",
			config: &config.Config{
				FilesFolder: "/test/folder",
			},
			file:           "",
			expectedResult: "/test/folder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := file.GetFilePath(tt.config, tt.file)
			if result != tt.expectedResult {
				t.Errorf("GetFilePath() = %q, want %q", result, tt.expectedResult)
			}
		})
	}
}

func TestGetProcessedFilePath(t *testing.T) {
	tests := []struct {
		name           string
		config         *config.Config
		file           string
		expectedResult string
	}{
		{
			name: "Simple processed file path",
			config: &config.Config{
				ProcessedFilesFolder: "/processed/folder",
			},
			file:           "file.mp4",
			expectedResult: "/processed/folder/file.mp4",
		},
		{
			name: "Processed file path with subdirectory",
			config: &config.Config{
				ProcessedFilesFolder: "/processed/folder",
			},
			file:           "subdir/file.mp4",
			expectedResult: "/processed/folder/subdir/file.mp4",
		},
		{
			name: "Path with trailing slash in processed config",
			config: &config.Config{
				ProcessedFilesFolder: "/processed/folder/",
			},
			file:           "file.mp4",
			expectedResult: "/processed/folder/file.mp4",
		},
		{
			name: "Path with relative processed file notation",
			config: &config.Config{
				ProcessedFilesFolder: "/processed/folder",
			},
			file:           "./file.mp4",
			expectedResult: "/processed/folder/file.mp4",
		},
		{
			name: "Empty processed file path",
			config: &config.Config{
				ProcessedFilesFolder: "/processed/folder",
			},
			file:           "",
			expectedResult: "/processed/folder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := file.GetProcessedFilePath(tt.config, tt.file)
			if result != tt.expectedResult {
				t.Errorf("GetProcessedFilePath() = %q, want %q", result, tt.expectedResult)
			}
		})
	}
}
