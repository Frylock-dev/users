//go:build wireinject
// +build wireinject

package injection

import (
	api "github.com/Frylock-dev/users/internal/api/user"
	repo "github.com/Frylock-dev/users/internal/repository/postgres/user"
	service "github.com/Frylock-dev/users/internal/service/user"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitUserAPI(conn *pgxpool.Pool) *api.API {
	wire.Build(
		repo.NewRepository,
		service.NewService,
		api.NewAPI,
	)

	return &api.API{}
}
