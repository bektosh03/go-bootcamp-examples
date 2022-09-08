package handler

import (
	"api-gateway/service"
)

func New(svc service.Service) Handler {
	return Handler{
		service: svc,
	}
}

type Handler struct {
	service service.Service
}
