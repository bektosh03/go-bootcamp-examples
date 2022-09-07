package httperr

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal server error")
)
