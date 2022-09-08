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
)

type JournalService struct {
	client journalpb.JournalServiceClient
}

func NewJournalService(client journalpb.JournalServiceClient) JournalService {
	return JournalService{
		client: client,
	}
}
func (s JournalService) RegisterJournal(ctx context.Context, req request.CreateJournalRequest) (response.Journal, error) {
	grpcRequest := &journalpb.CreateJournalRequest{
		ScheduleId: req.ScheduleID,
		StudentId:  req.StudentID,
		Attended:   req.Attended,
		Mark:       req.Mark,
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
		StudentID:  res.StudentId,
		Attended:   res.Attended,
		Mark:       res.Mark,
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
		StudentID:  res.StudentId,
		Attended:   false,
		Mark:       0,
	}, nil
}
func (s JournalService) UpdateJournal(ctx context.Context, req request.Journal) (response.Journal, error) {
	res, err := s.client.UpdateJournal(ctx, &journalpb.Journal{
		Id:         req.ID,
		ScheduleId: req.ScheduleID,
		StudentId:  req.StudentID,
		Attended:   req.Attended,
		Mark:       req.Mark,
	})
	if err != nil {
		return response.Journal{}, err
	}
	return response.Journal{
		ID:         res.Id,
		ScheduleID: res.ScheduleId,
		StudentID:  res.StudentId,
		Attended:   res.Attended,
		Mark:       res.Mark,
	}, nil
}
func (s JournalService) DeleteJournal(ctx context.Context, id string) error {
	_, err := s.client.DeleteJournal(ctx, &journalpb.DeleteJournalRequest{JournalId: id})
	return err
}
