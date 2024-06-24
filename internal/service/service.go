package service

import (
	"github.com/Frylock-dev/users/internal/model"
	"golang.org/x/net/context"
)

type User interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (string, error)
	GetByUUID(ctx context.Context, uuid string) (*model.User, error)
}
