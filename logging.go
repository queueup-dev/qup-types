package types

type Logger interface {
	Error(content DataNode, err error, message *string)
	Warning(content DataNode, message string, err error)
	Info(content DataNode, message string)
	Debug(content DataNode, message string, err error)
}
