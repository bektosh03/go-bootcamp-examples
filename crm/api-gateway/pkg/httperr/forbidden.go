package httperr

import (
	"github.com/go-chi/render"
	"net/http"
)

func Forbidden(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, render.M{
		"error": ErrForbidden.Error(),
		"msg":   msg,
	})
}
