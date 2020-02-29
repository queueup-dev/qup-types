package types

type Logger interface {
	Error(err error, message *string)
	Warning(message string, err error)
	Info(message string)
	Debug(message string, err error)
}
