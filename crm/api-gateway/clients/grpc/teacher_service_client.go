package grpc

import (
	"context"

	"github.com/bektosh03/crmprotos/teacherpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTeacherServiceClient(ctx context.Context, url string) (teacherpb.TeacherServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx, url,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return teacherpb.NewTeacherServiceClient(conn), nil
}
