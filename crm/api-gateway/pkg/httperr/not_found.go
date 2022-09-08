package httperr

import (
	"net/http"

	"github.com/go-chi/render"
)

func NotFoundErr(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, render.M{
		"error":   ErrNotFound.Error(),
		"message": err,
	})
}
