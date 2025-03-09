package logger

type Logger interface {
	Debug(msg string, values ...interface{})
	Info(msg string, values ...interface{})
	Error(msg string, values ...interface{})
	Fatal(msg string, values ...interface{})
}

func NewLogger(serviceName string) (Logger, error) {
	return newZapLogger(serviceName)
}
