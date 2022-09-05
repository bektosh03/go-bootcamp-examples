package server

import (
	"context"
	"schedule-service/service"

	"github.com/bektosh03/crmprotos/schedulepb"
)

type Server struct {
	schedulepb.UnimplementedScheduleServiceServer
	svc service.Service
}

func New(svc service.Service) Server {
	return Server{
		svc: svc,
	}
}

func (s Server) CreateSchedule(ctx context.Context, req *schedulepb.CreateScheduleRequest) (*schedulepb.Schedule, error) {

}
