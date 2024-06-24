package repository

import (
	"context"
	"github.com/Frylock-dev/users/internal/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=User
type User interface {
	Save(ctx context.Context, uuid string, userInfo *model.UserInfo) error
	GetByUUID(ctx context.Context, uuid string) (*model.User, error)
}
