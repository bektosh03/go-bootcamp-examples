package main

import (
	"github.com/Shopify/sarama"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmprotos/journalpb"
	"google.golang.org/grpc"
	"journal-service/config"
	consumer2 "journal-service/consumer"
	"journal-service/domain/journal"
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

	kafkaCfg := sarama.NewConfig()
	client, err := sarama.NewClient(
		[]string{net.JoinHostPort(cfg.KafkaHost, cfg.KafkaPort)},
		kafkaCfg,
	)
	if err != nil {
		panic(err)
	}
	consumer, err := consumer2.NewConsumer(client)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	repo, err := repository.NewPostgresRepository(cfg.Config)
	if err != nil {
		panic(err)
	}

	journalFactory := journal.NewFactory(id.Generator{})
	svc := service.New(repo, journalFactory, consumer)
	svr := server.New(svc, journalFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	journalpb.RegisterJournalServiceServer(grpcServer, svr)

	svc.Run()

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
