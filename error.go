package basm

import "errors"

// errWrap avoids the use of the fmt package to wrap errors.
func errWrap(msg string, err error) error {
	e := errors.New(msg)
	return errors.Join(e, err)
}
