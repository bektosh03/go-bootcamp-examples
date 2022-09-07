package server

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"journal-service/domain/journal"
	journalpb "journal-service/protos"
	"journal-service/service"
)

type Server struct {
	journalpb.UnimplementedJournalServiceServer
	service        service.Service
	journalFactory journal.Factory
}

func New(svc service.Service, journalFactory journal.Factory) Server {
	return Server{
		service:        service.Service{},
		journalFactory: journal.Factory{},
	}
}

func (s Server) RegisterJournal(ctx context.Context, req *journalpb.RegisterJournalRequest) (*journalpb.Journal, error) {
	jour, err := s.convertRegisterJournalRequestToDomainJournal(req)
	if err != nil {
		return nil, err
	}
	if err = s.service.RegisterJournal(ctx, jour); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return toProtoJournal(jour), nil
}

func (s Server) convertRegisterJournalRequestToDomainJournal(protoJour *journalpb.RegisterJournalRequest) (journal.Journal, error) {
	scheduleId, err := uuid.Parse(protoJour.ScheduleId)
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "provided schedule id is not uuid")
	}
	studentId, err := uuid.Parse(protoJour.ScheduleId)
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "provided student id is not uuid")
	}

	jour, err := s.journalFactory.NewJournal(
		scheduleId, studentId, protoJour.Attended, protoJour.Mark,
	)
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return jour, nil

}
