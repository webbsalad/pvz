package app

import (
	"github.com/webbsalad/pvz/internal/config"
	"go.uber.org/fx"

	item_api "github.com/webbsalad/pvz/internal/api/item"
	login_api "github.com/webbsalad/pvz/internal/api/login"
	pvz_api "github.com/webbsalad/pvz/internal/api/pvz"
	pb "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"

	item_service "github.com/webbsalad/pvz/internal/service/item/v1"
	login_service "github.com/webbsalad/pvz/internal/service/login/v1"
	pvz_service "github.com/webbsalad/pvz/internal/service/pvz/v1"

	item_repository "github.com/webbsalad/pvz/internal/repository/item/pg"
	login_repository "github.com/webbsalad/pvz/internal/repository/login/pg"
	pvz_repository "github.com/webbsalad/pvz/internal/repository/pvz/pg"
	user_repository "github.com/webbsalad/pvz/internal/repository/user/pg"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			initDB,
		),

		grpcOption(),
		gatewayOption(),

		// services
		servicesOption(),
	)
}

func servicesOption() fx.Option {
	return fx.Options(
		fx.Provide(
			login_api.NewImplementation,
			pvz_api.NewImplementation,
			item_api.NewImplementation,

			login_service.NewService,
			pvz_service.NewService,
			item_service.NewService,

			user_repository.NewRepository,
			login_repository.NewRepository,
			pvz_repository.NewRepository,
			item_repository.NewRepository,
		),
		fx.Invoke(
			pb.RegisterPVZServiceServer,
			pb.RegisterItemServiceServer,
			pb.RegisterLoginServiceServer,
		),
	)
}
