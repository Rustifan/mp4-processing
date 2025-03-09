package nats

type ProcessFileDto struct {
	FilePath string `json:"filePath"`
}

type FileUpdateTopicDto struct {
	FilePath         string `json:"filePath"`
	Status           string `json:"status"`
	ProcssedFilePath string `json:"procssedFilePath"`
}
