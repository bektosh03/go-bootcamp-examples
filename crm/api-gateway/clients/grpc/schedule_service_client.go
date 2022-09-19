package grpc

import (
	"context"
	"github.com/bektosh03/crmprotos/schedulepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewScheduleServiceClient(ctx context.Context, url string) (schedulepb.ScheduleServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx, url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return schedulepb.NewScheduleServiceClient(conn), nil
}
