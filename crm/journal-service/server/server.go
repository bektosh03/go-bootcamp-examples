package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/bektosh03/crmcommon/errs"
	"journal-service/domain/journal"
	"journal-service/service"
	"time"

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
		service:        svc,
		journalFactory: journalFactory,
	}
}

func (s Server) CreateJournal(ctx context.Context, req *journalpb.CreateJournalRequest) (*journalpb.Journal, error) {
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

func (s Server) GetJournal(ctx context.Context, req *journalpb.GetJournalRequest) (*journalpb.Journal, error) {
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

func (s Server) MarkStudent(ctx context.Context, req *journalpb.MarkStudentRequest) (*emptypb.Empty, error) {
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	journalID, err := uuid.Parse(req.GetJournalId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "journal id is not uuid")
	}

	st, err := journal.NewStatus(journalID, studentID, true, req.Mark)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.service.MarkStudent(ctx, st); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s Server) SetStudentAttendance(ctx context.Context, req *journalpb.SetStudentAttendanceRequest) (*emptypb.Empty, error) {
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	journalID, err := uuid.Parse(req.GetJournalId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "journal id is not uuid")
	}

	st, err := journal.NewStatus(journalID, studentID, req.Attended, 0)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.service.SetStudentAttendance(ctx, st); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s Server) GetStudentJournalEntries(ctx context.Context, req *journalpb.GetStudentJournalEntriesRequest) (*journalpb.Entries, error) {
	studentID, err := uuid.Parse(req.StudentId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	entries, err := s.service.GetStudentJournalEntries(ctx, studentID, req.TimeRange.Start.AsTime(), req.TimeRange.End.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoEntries(entries), nil
}

func (s Server) GetTeacherJournalEntries(ctx context.Context, req *journalpb.GetTeacherJournalEntriesRequest) (*journalpb.Entries, error) {
	teacherID, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	entries, err := s.service.GetTeacherJournalEntries(ctx, teacherID, req.TimeRange.Start.AsTime(), req.TimeRange.End.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoEntries(entries), nil

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

	teacherID, err := uuid.Parse(protoJour.TeacherId)
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, "provided teacher id is not uuid")
	}

	jour, err := s.journalFactory.NewJournal(scheduleId, teacherID, protoJour.Date.AsTime())
	if err != nil {
		return journal.Journal{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return jour, nil
}
