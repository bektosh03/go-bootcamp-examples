package main

import (
	"fmt"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmprotos/studentpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"student-service/config"
	"student-service/domain/group"
	"student-service/domain/student"
	"student-service/repository"
	"student-service/server"
	"student-service/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("error with loading config: ", err)
	}

	repo, err := repository.NewPostgres(cfg.Config)
	if err != nil {
		panic(err)
	}

	groupFactory := group.NewFactory(id.Generator{})
	studentFactory := student.NewFactory(id.Generator{})

	svc := service.New(repo, groupFactory, studentFactory)
	svr := server.New(svc, groupFactory, studentFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, svr)

	fmt.Println("Server starting at:", lis.Addr().String())

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
