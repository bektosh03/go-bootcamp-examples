package handler

import (
	"api-gateway/pkg/producer"
	"api-gateway/service"
)

func New(svc service.Service, producer producer.Producer) Handler {
	return Handler{
		service:  svc,
		producer: producer,
	}
}

type Handler struct {
	service  service.Service
	producer producer.Producer
}
