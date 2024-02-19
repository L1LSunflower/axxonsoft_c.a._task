package cerrors

import "errors"

var (
	ErrReadBody = errors.New("can't read body")
)
