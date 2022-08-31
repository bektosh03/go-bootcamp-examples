package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"api-gateway/adapter"
	"api-gateway/clients/grpc"
	"api-gateway/handler"
	"api-gateway/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const TeacherServiceURL = "localhost:8001"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	teacherServiceClient, err := grpc.NewTeacherServiceClient(ctx, TeacherServiceURL)
	if err != nil {
		log.Panicln("failed to create new teacher service client:", err)
	}

	teacherService := adapter.NewTeacherService(teacherServiceClient)
	service := service.New(teacherService)
	h := handler.New(service)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/teacher", h.RegisterTeacher)
	r.Get("/teacher/{id}", h.GetTeacher)
	r.Post("/subject", h.CreateSubject)
	r.Get("/subject/{id}", h.GetSubject)

	http.ListenAndServe(":8080", r)
}
