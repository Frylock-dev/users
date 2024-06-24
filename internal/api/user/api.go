package user

import (
	"github.com/Frylock-dev/users/internal/converter"
	"github.com/Frylock-dev/users/internal/service"
	"github.com/Frylock-dev/users/pkg/user_v1"
	"golang.org/x/net/context"
)

type API struct {
	user_v1.UnimplementedUserServiceServer
	userService service.User
}

func NewAPI(userService service.User) *API {
	return &API{userService: userService}
}

func (api *API) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	uuid, err := api.userService.Create(ctx, converter.FromUserInfoToService(req.Info))
	if err != nil {
		return nil, err
	}

	return &user_v1.CreateResponse{Uuid: uuid}, nil
}

func (api *API) GetByUUID(ctx context.Context, req *user_v1.GetByUUIDRequest) (*user_v1.GetByUUIDResponse, error) {
	user, err := api.userService.GetByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &user_v1.GetByUUIDResponse{User: converter.FromUserToGRPC(user)}, nil
}
