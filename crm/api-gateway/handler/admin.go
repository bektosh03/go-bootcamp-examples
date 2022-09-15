package handler

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"net/http"

	"github.com/go-chi/render"
)

func (h Handler) AuthAdmin(w http.ResponseWriter, r *http.Request) {
	var req request.AdminRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	if req.Password != auth.AdminPassword || req.UserName != auth.AdminUserName {
		httperr.Unauthorized(w, r, "password or user_name for admin is incorrect")
		return
	}

	render.JSON(w, r, render.M{
		"ok":    true,
		"token": auth.NewJWTForAdmin(),
	})

}
