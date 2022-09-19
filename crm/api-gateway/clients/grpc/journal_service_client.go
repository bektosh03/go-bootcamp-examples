package grpc

import (
	"context"
	"github.com/bektosh03/crmprotos/journalpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewJournalServiceClient(ctx context.Context, url string) (journalpb.JournalServiceClient, error) {
	conn, err := grpc.DialContext(
		ctx, url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return journalpb.NewJournalServiceClient(conn), nil
}
