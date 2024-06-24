package main

import (
	"fmt"
	"github.com/Frylock-dev/users/config"
	"github.com/Frylock-dev/users/internal/injection"
	"github.com/Frylock-dev/users/internal/interceptor"
	"github.com/Frylock-dev/users/internal/metric"
	"github.com/Frylock-dev/users/internal/tracing"
	"github.com/Frylock-dev/users/pkg/user_v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	appName   = "users"
	namespace = "dc"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = metric.Init(ctx, appName, namespace)
	if err != nil {
		log.Fatal(err)
	}

	tracing.Init(appName)

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			interceptor.MetricsInterceptor,
			interceptor.LogInterceptor,
			interceptor.ServerTracingInterceptor,
		),
	))

	reflection.Register(srv)

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	api := injection.InitUserAPI(pool)

	user_v1.RegisterUserServiceServer(srv, api)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port %d", cfg.Port)
	log.Fatal(srv.Serve(listen))
}
