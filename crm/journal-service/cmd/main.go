package main

import (
	"github.com/bektosh03/crmcommon/id"
	"google.golang.org/grpc"
	"journal-service/config"
	"journal-service/domain/journal"
	journalpb "journal-service/protos"
	"journal-service/repository"
	"journal-service/server"
	"journal-service/service"
	"log"
	"net"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("error with loading config: ", err)
	}
	repo, err := repository.NewPostgresRepository(cfg.Config)
	if err != nil {
		panic(err)
	}

	journalFactory := journal.NewFactory(id.Generator{})
	svc := service.New(repo, journalFactory)
	svr := server.New(svc, journalFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	journalpb.RegisterJournalServiceServer(grpcServer, svr)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
