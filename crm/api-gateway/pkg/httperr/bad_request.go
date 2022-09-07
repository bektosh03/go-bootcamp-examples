package httperr

import (
	"net/http"

	"github.com/go-chi/render"
)

func InvalidJSON(w http.ResponseWriter, r *http.Request) {
	BadRequest(w, r, "invalid json")
}

func BadRequest(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, render.M{
		"error":   ErrBadRequest.Error(),
		"message": msg,
	})
}
