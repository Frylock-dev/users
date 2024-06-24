package user

import (
	"github.com/Frylock-dev/users/internal/model"
	"github.com/Frylock-dev/users/internal/repository"
	"github.com/Frylock-dev/users/internal/service"
	goouid "github.com/google/uuid"
	"golang.org/x/net/context"
)

type Service struct {
	userRepo repository.User
}

func NewService(userRepo repository.User) service.User {
	return &Service{userRepo: userRepo}
}

func (s *Service) Create(ctx context.Context, userInfo *model.UserInfo) (string, error) {
	uuid := goouid.New().String()

	err := s.userRepo.Save(ctx, uuid, userInfo)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (s *Service) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	user, err := s.userRepo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
