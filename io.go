package types

import "io"

type Validator interface {
	Valid() bool
}

type PayloadWriter interface {
	Validator
	Marshal() ([]byte, error)
	Reader() (io.Reader, error)
	ToString() (*string, error)
	GetTarget() string
}

type PayloadReader interface {
	Validator
	Bytes() ([]byte, error)
	Unmarshal(structure interface{}) error
	Reader() io.Reader
	ToString() (*string, error)
	GetSource() string
}
