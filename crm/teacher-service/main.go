package main

import (
	"fmt"
	"net"
	"teacher-service/config"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/pkg/id"
	"teacher-service/repository"
	"teacher-service/server"
	"teacher-service/service"

	"github.com/bektosh03/crmprotos/teacherpb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	repo, err := repository.NewPostgres(cfg.Config)
	if err != nil {
		panic(err)
	}

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	svc := service.New(repo, subjectFactory, teacherFactory)
	server := server.New(svc, subjectFactory, teacherFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	teacherpb.RegisterTeacherServiceServer(grpcServer, server)

	fmt.Println("Server starting at:", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
