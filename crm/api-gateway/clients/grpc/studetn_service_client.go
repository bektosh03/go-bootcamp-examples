package grpc

import (
	"context"
	"github.com/bektosh03/crmprotos/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStudentServiceClient(ctx context.Context, url string) (studentpb.StudentServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx, url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return studentpb.NewStudentServiceClient(conn), nil
}
