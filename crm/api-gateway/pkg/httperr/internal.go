package httperr

import (
	"net/http"

	"github.com/go-chi/render"
)

func InternalError(w http.ResponseWriter, r *http.Request, messag string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, render.M{
		"error":   ErrInternal.Error(),
		"message": messag,
	})
}
