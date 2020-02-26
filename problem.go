package types

type ProblemResponse interface {
	error
	StatusCode() int16
	ParseBody() string
}
