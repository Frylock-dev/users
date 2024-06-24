package user

import (
	"errors"
	"github.com/Frylock-dev/users/internal/model/stub"
	"github.com/Frylock-dev/users/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"testing"
)

func TestService_Create_Positive(t *testing.T) {
	ctx := context.Background()
	userRepoMock := mocks.NewUser(t)
	service := NewService(userRepoMock)
	userStub := stub.NewUserStub()

	userRepoMock.On(
		"Save",
		ctx,
		mock.Anything,
		userStub.Info,
	).Return(nil)

	uuid, err := service.Create(ctx, userStub.Info)
	assert.NoError(t, err)
	assert.NotEmpty(t, uuid)
}

func TestService_Create_With_Repo_Error(t *testing.T) {
	ctx := context.Background()
	userRepoMock := mocks.NewUser(t)
	service := NewService(userRepoMock)
	userStub := stub.NewUserStub()

	userRepoMock.On(
		"Save",
		ctx,
		mock.Anything,
		userStub.Info,
	).
		Return(errors.New("repo error"))

	_, err := service.Create(ctx, userStub.Info)
	assert.Error(t, err)
}

func TestService_GetByUUID_Positive(t *testing.T) {
	ctx := context.Background()
	userRepoMock := mocks.NewUser(t)
	service := NewService(userRepoMock)
	userStub := stub.NewUserStub()

	userRepoMock.
		On("GetByUUID", ctx, userStub.UUID).
		Return(userStub, nil)

	user, err := service.GetByUUID(ctx, userStub.UUID)
	assert.NoError(t, err)
	assert.Equal(t, userStub.UUID, user.UUID)
}

func TestService_GetByUUID_With_Repo_Error(t *testing.T) {
	ctx := context.Background()
	userRepoMock := mocks.NewUser(t)
	service := NewService(userRepoMock)
	userStub := stub.NewUserStub()

	userRepoMock.
		On("GetByUUID", ctx, userStub.UUID).
		Return(nil, errors.New("repo error"))

	_, err := service.GetByUUID(ctx, userStub.UUID)
	assert.Error(t, err)
}
