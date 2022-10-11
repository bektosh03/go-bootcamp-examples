package main

import (
	"github.com/Shopify/sarama"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmprotos/journalpb"
	"google.golang.org/grpc"
	"journal-service/config"
	"journal-service/consumer/attend"
	"journal-service/consumer/mark"
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
		log.Println("error creating sarama NewClient",err)
		return
	}

	attConsumer,err:=attend.NewConsumer(client)
	if err != nil {
		log.Println("error creating new attConsumer",err)
		return
	}

	markConsumer,err:=mark.NewConsumer(client)
	if err != nil {
		log.Println("error creating new markConsumer",err)
		return
	}

	defer client.Close()

	repo, err := repository.NewPostgresRepository(cfg.Config)
	if err != nil {
		log.Println("error creating new postgres repo",err)
		return
	}

	journalFactory := journal.NewFactory(id.Generator{})
	svc := service.New(repo, journalFactory, attConsumer,markConsumer)
	svr := server.New(svc, journalFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		log.Println("error with listening",err)
		return
	}
	grpcServer := grpc.NewServer()
	journalpb.RegisterJournalServiceServer(grpcServer, svr)

	svc.RunAtt()
	svc.RunMark()

	if err = grpcServer.Serve(lis); err != nil {
		log.Println("error with serving grpc",err)
	}
}
