package basm

import "errors"

func errWrap(msg string, err error) error {
	e := errors.New(msg)
	return errors.Join(e, err)
}
