package main

import (
	"api-gateway/config"
	"api-gateway/pkg/auth"
	"api-gateway/pkg/producer"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/cors"

	"api-gateway/adapter"
	"api-gateway/clients/grpc"
	"api-gateway/handler"
	"api-gateway/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	StudentServiceURL  = "localhost:8002"
	TeacherServiceURL  = "localhost:8001"
	ScheduleServiceURL = "localhost:8003"
	JournalServiceURL  = "localhost:8004"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Panicln("failed to load config", err)
	}

	fmt.Println("cfg:", cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	teacherServiceClient, err := grpc.NewTeacherServiceClient(ctx, cfg.TeacherServiceAddr)
	if err != nil {
		log.Panicln("failed to create new teacher service client:", err)
	}

	studentServiceClient, err := grpc.NewStudentServiceClient(ctx, cfg.StudentServiceAddr)
	if err != nil {
		log.Panicln("failed to create new student service client:", err)
	}
	scheduleServiceClient, err := grpc.NewScheduleServiceClient(ctx, cfg.ScheduleServiceAddr)
	if err != nil {
		log.Panicln("failed to create new schedule service client:", err)
	}

	journalServiceClient, err := grpc.NewJournalServiceClient(ctx, cfg.JournalServiceAddr)
	if err != nil {
		log.Panicln("failed to create new journal service client:", err)
	}
	teacherService := adapter.NewTeacherService(teacherServiceClient)
	studentService := adapter.NewStudentService(studentServiceClient)
	scheduleService := adapter.NewScheduleService(scheduleServiceClient)
	journalService := adapter.NewJournalService(journalServiceClient)

	svc := service.New(teacherService, studentService, scheduleService, journalService)

	kafkaCfg := sarama.NewConfig()
	kafkaCfg.Producer.Return.Successes = true
	kafkaCfg.Producer.Return.Errors = true

	fmt.Println("KAFKA:", net.JoinHostPort(cfg.KafkaHost, cfg.KafkaPort))

	kafkaClient, err := sarama.NewClient(
		[]string{net.JoinHostPort(cfg.KafkaHost, cfg.KafkaPort)},
		kafkaCfg,
	)
	if err != nil {
		log.Panicln("failed to initialize kafka client:", err)
	}
	kafkaProducer, err := producer.NewKafkaProducer(kafkaClient)
	if err != nil {
		log.Panicln("failed to initialize kafka producer:", err)
	}

	h := handler.New(svc, kafkaProducer)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)

	r.Post("/admin", h.AuthAdmin)
	// registration endpoints
	r.Group(func(r chi.Router) {
		r.Use(auth.AdminMiddleware)
		r.Post("/teacher", h.RegisterTeacher)
		r.Post("/student", h.RegisterStudent)
		r.Post("/subject", h.CreateSubject)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Post("/journal", h.RegisterJournal)
		r.Get("/journal/{id}", h.GetJournal)
		r.Put("/journal", h.UpdateJournal)
		r.Delete("/journal/{id}", h.DeleteJournal)
		r.Put("/journal/attendance", h.SetStudentAttendance)
		r.Put("/journal/mark", h.MarkStudent)
		r.Get("/journal/student/{id}", h.GetStudentJournal)
		r.Get("/journal/teacher/{id}", h.GetTeacherJournal)
	})
	// teacher endpoints
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Post("/get/teacher", h.GetTeacher)
		r.Delete("/teacher/delete/{id}", h.DeleteTeacher)
		r.Get("/teachers", h.ListTeachers)
	})

	// subject endpoints
	r.Group(func(r chi.Router) {
		// TODO add auth middleware
		r.Get("/subject/{id}", h.GetSubject)
		r.Delete("/subject/delete/{id}", h.DeleteSubject)
		r.Get("/subjects", h.ListSubjects)
	})

	// student endpoints
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Post("/get/student", h.GetStudent)
		r.Put("/student", h.UpdateStudent)
		r.Delete("/student/{id}", h.DeleteStudent)
		r.Get("/students", h.ListStudents)
	})

	// group endpoints
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Post("/group", h.CreateGroup)
		r.Get("/group/{id}", h.GetGroup)
		r.Put("/group", h.UpdateGroup)
		r.Delete("/group/{id}", h.DeleteGroup)
		r.Get("/groups", h.ListGroups)
	})

	// schedule endpoints
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Post("/schedule", h.RegisterSchedule)
		r.Get("/schedule/{id}", h.GetScheduleById)
		r.Get("/schedule/teacher/{id}", h.GetFullScheduleForTeacher)
		r.Get("/schedule/group/{id}", h.GetFullScheduleForGroup)
		r.Post("/schedule/teacher", h.GetSpecificDateScheduleForTeacher)
		r.Post("/schedule/group", h.GetSpecificDateScheduleForGroup)
	})

	fmt.Println("Server starting at :8080")

	http.ListenAndServe(":8080", r)
}
