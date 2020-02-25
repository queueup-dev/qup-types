package symbols

import "io"

type Validator interface {
	Valid() bool
}

type PayloadWriter interface {
	Validator
	Marshal() ([]byte, error)
	ToString() (*string, error)
}

type PayloadReader interface {
	Validator
	Bytes() ([]byte, error)
	Unmarshal(structure interface{}) error
	Reader() io.Reader
	ToString() (*string, error)
}
