package httperr

import (
	"github.com/go-chi/render"
	"net/http"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, render.M{
		"error": ErrUnauthorized.Error(),
		"msg":   msg,
	})
}
