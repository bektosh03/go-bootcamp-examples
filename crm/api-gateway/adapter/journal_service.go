package adapter

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"fmt"
	"github.com/bektosh03/crmprotos/journalpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type JournalService struct {
	client journalpb.JournalServiceClient
}

func NewJournalService(client journalpb.JournalServiceClient) JournalService {
	return JournalService{
		client: client,
	}
}

func (s JournalService) GetStudentJournal(ctx context.Context, studentID string, start, end time.Time) ([]response.JournalEntry, error) {
	res, err := s.client.GetStudentJournalEntries(ctx, &journalpb.GetStudentJournalEntriesRequest{
		StudentId: studentID,
		TimeRange: &journalpb.TimeRange{
			Start: &timestamppb.Timestamp{
				Seconds: start.Unix(),
				Nanos:   int32(start.UnixNano()),
			},
			End: &timestamppb.Timestamp{
				Seconds: end.Unix(),
				Nanos:   int32(end.UnixNano()),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return fromProtoToResponseEntries(res), nil
}

func (s JournalService) MarkStudent(ctx context.Context, req request.MarkStudentRequest) error {
	_, err := s.client.MarkStudent(ctx, &journalpb.MarkStudentRequest{
		Mark:      req.Mark,
		StudentId: req.StudentID,
		JournalId: req.JournalID,
	})

	return err
}

func (s JournalService) SetStudentAttendance(ctx context.Context, req request.SetStudentAttendanceRequest) error {
	_, err := s.client.SetStudentAttendance(ctx, &journalpb.SetStudentAttendanceRequest{
		Attended:  req.Attended,
		StudentId: req.StudentID,
		JournalId: req.JournalID,
	})

	return err
}

func (s JournalService) RegisterJournal(ctx context.Context, scheduleID string, date time.Time, studentIDs []string) (response.Journal, error) {
	grpcRequest := &journalpb.CreateJournalRequest{
		ScheduleId: scheduleID,
		Date: &timestamppb.Timestamp{
			Seconds: date.Unix(),
			Nanos:   int32(date.UnixNano()),
		},
		StudentIds: studentIDs,
	}
	res, err := s.client.CreateJournal(ctx, grpcRequest)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return response.Journal{}, fmt.Errorf(" %w: %s", httperr.ErrBadRequest, st.Message())
			default:
				return response.Journal{}, fmt.Errorf("%w: %s", httperr.ErrInternal, st.Message())
			}
		}
		return response.Journal{}, err
	}

	return response.Journal{
		ID:         res.Id,
		ScheduleID: res.ScheduleId,
		Date:       res.Date.AsTime(),
	}, nil
}
func (s JournalService) GetJournal(ctx context.Context, journalId string) (response.Journal, error) {
	res, err := s.client.GetJournal(ctx, &journalpb.GetJournalRequest{JournalId: journalId})
	if err != nil {
		return response.Journal{}, err
	}
	return response.Journal{
		ID:         res.Id,
		ScheduleID: res.ScheduleId,
		Date:       res.Date.AsTime(),
	}, nil
}
func (s JournalService) UpdateJournal(ctx context.Context, req request.Journal) (response.Journal, error) {
	res, err := s.client.UpdateJournal(ctx, &journalpb.Journal{
		Id:         req.ID,
		ScheduleId: req.ScheduleID,
		Date: &timestamppb.Timestamp{
			Seconds: req.Date.Unix(),
			Nanos:   int32(req.Date.UnixNano()),
		},
	})
	if err != nil {
		return response.Journal{}, err
	}
	return response.Journal{
		ID:         res.Id,
		ScheduleID: res.ScheduleId,
		Date:       res.Date.AsTime(),
	}, nil
}
func (s JournalService) DeleteJournal(ctx context.Context, id string) error {
	_, err := s.client.DeleteJournal(ctx, &journalpb.DeleteJournalRequest{JournalId: id})
	return err
}
