package server

import (
	"context"
	"errors"
	"fmt"
	"journal-service/domain/journal"
	"journal-service/service"
	"time"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmprotos/journalpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s Server) RegisterJournal(ctx context.Context, req *journalpb.CreateJournalRequest) (*journalpb.Journal, error) {
	jour, err := s.convertRegisterJournalRequestToDomainJournal(req)
	if err != nil {
		return nil, err
	}

	studentIDs := make([]uuid.UUID, 0, len(req.StudentIds))
	for _, idString := range req.StudentIds {
		id, err := uuid.Parse(idString)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("student id=%q is not uuid", idString))
		}
		studentIDs = append(studentIDs, id)
	}

	if err = s.service.RegisterJournal(ctx, jour, studentIDs); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoJournal(jour), nil
}
func (s Server) GetJournalEntry(ctx context.Context, req *journalpb.GetJournalEntryRequest) (*journalpb.Journal, error) {
	id, err := uuid.Parse(req.JournalId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}
	jour, err := s.service.GetJournal(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return toProtoJournal(jour), nil
}
func (s Server) UpdateJournal(ctx context.Context, req *journalpb.Journal) (*journalpb.Journal, error) {
	jour, err := s.convertUpdateJournalRequestToDomainJournal(req)
	if err != nil {
		return nil, err
	}
	if err = s.service.UpdateJournal(ctx, jour); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return toProtoJournal(jour), err
}
func (s Server) DeleteJournal(ctx context.Context, req *journalpb.DeleteJournalRequest) (*emptypb.Empty, error) {
	journalId, err := uuid.Parse(req.GetJournalId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "journal id is not uuid")
	}
	if err = s.service.DeleteJournal(ctx, journalId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s Server) convertUpdateJournalRequestToDomainJournal(protojour *journalpb.Journal) (journal.Journal, error) {
	id, err := uuid.Parse(protojour.GetId())
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "journal id is not uuid")
	}
	scheduleId, err := uuid.Parse(protojour.GetScheduleId())
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "schedule id is not uuid")
	}
	jour, err := journal.UnmarshalJournal(journal.UnmarshalJournalArgs{
		ID:         id,
		ScheduleID: scheduleId,
		Date:       time.Time{},
	})
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return jour, nil
}

func (s Server) convertRegisterJournalRequestToDomainJournal(protoJour *journalpb.CreateJournalRequest) (journal.Journal, error) {
	scheduleId, err := uuid.Parse(protoJour.ScheduleId)
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "provided schedule id is not uuid")
	}

	jour, err := s.journalFactory.NewJournal(scheduleId, protoJour.Date.AsTime())
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return jour, nil
}
