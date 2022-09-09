package group

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateGroup(ctx context.Context, s Group) error
	GetGroup(ctx context.Context, id uuid.UUID) (Group, error)
	UpdateGroup(ctx context.Context, s Group) error
	DeleteGroup(ctx context.Context, id uuid.UUID) error
	ListGroups(ctx context.Context, page, limit int32) ([]Group, int, error)
}
