package main

import (
	"context"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"time"

	"api-gateway/adapter"
	"api-gateway/clients/grpc"
	"api-gateway/handler"
	"api-gateway/service"

	"github.com/go-chi/chi/v5"
)

const (
	StudentServiceURL  = "localhost:8002"
	TeacherServiceURL  = "localhost:8001"
	ScheduleServiceURL = "localhost:8003"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	teacherServiceClient, err := grpc.NewTeacherServiceClient(ctx, TeacherServiceURL)
	if err != nil {
		log.Panicln("failed to create new teacher service client:", err)
	}

	studentServiceClient, err := grpc.NewStudentServiceClient(ctx, StudentServiceURL)
	if err != nil {
		log.Panicln("failed to create new student service client:", err)
	}
	scheduleServiceClient, err := grpc.NewScheduleServiceClient(ctx, ScheduleServiceURL)
	if err != nil {
		log.Panicln("failed to create new schedule service client:", err)
	}

	teacherService := adapter.NewTeacherService(teacherServiceClient)
	studentService := adapter.NewStudentService(studentServiceClient)
	scheduleService := adapter.NewScheduleService(scheduleServiceClient)

	service := service.New(teacherService, studentService, scheduleService)
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

	r.Post("/student", h.RegisterStudent)
	r.Get("/student/{id}", h.GetStudent)
	r.Put("/student", h.UpdateStudent)
	r.Delete("/student/{id}", h.DeleteStudent)
	r.Get("/students", h.ListStudents)
	r.Post("/group", h.CreateGroup)
	r.Get("/group/{id}", h.GetGroup)
	r.Put("/group", h.UpdateGroup)
	r.Delete("/group/{id}", h.DeleteGroup)
	r.Get("/groups", h.ListGroups)

	r.Post("/schedule", h.RegisterSchedule)
	r.Get("/schedule/{id}", h.GetScheduleById)

	http.ListenAndServe(":8080", r)
}
