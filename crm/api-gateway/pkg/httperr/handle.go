package httperr

import (
	"errors"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrBadRequest):
		BadRequest(w, r, err.Error())
	default:
		InternalError(w, r, err.Error())
	}
}
