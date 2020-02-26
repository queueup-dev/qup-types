package types

import "io"

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
