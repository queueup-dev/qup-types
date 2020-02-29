package types

type Logger interface {
	Error(err error, message *string)
	Warning(message string)
	Info(message string)
	Debug(message string)
}
