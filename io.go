package types

import (
	"github.com/getsentry/sentry-go"
	"io"
)

type DataNode interface {
	ApplyToSentryEvent(*sentry.Event)
	GetPayload() io.Reader
}

type Validator interface {
	Valid() bool
}

type PayloadVertex interface {
	Bytes() ([]byte, error)
	Reader() (io.Reader, error)
	ToString() (*string, error)
}

type PayloadWriter interface {
	Validator
	PayloadVertex
	Marshal() (interface{}, error)
}

type PayloadReader interface {
	Validator
	PayloadVertex
	Unmarshal(interface{}) error
}
