package app

import (
	"github.com/webbsalad/pvz/internal/config"
	"go.uber.org/fx"

	login_api "github.com/webbsalad/pvz/internal/api/login"
	pvz_api "github.com/webbsalad/pvz/internal/api/pvz"
	reception_api "github.com/webbsalad/pvz/internal/api/reception"
	pb "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"

	login_service "github.com/webbsalad/pvz/internal/service/login/v1"
	pvz_service "github.com/webbsalad/pvz/internal/service/pvz/v1"
	reception_service "github.com/webbsalad/pvz/internal/service/reception/v1"

	login_repository "github.com/webbsalad/pvz/internal/repository/login/pg"
	pvz_repository "github.com/webbsalad/pvz/internal/repository/pvz/pg"
	reception_repository "github.com/webbsalad/pvz/internal/repository/reception/pg"
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
			reception_api.NewImplementation,

			login_service.NewService,
			pvz_service.NewService,
			reception_service.NewService,

			user_repository.NewRepository,
			login_repository.NewRepository,
			pvz_repository.NewRepository,
			reception_repository.NewRepository,
		),
		fx.Invoke(
			pb.RegisterPVZServiceServer,
			pb.RegisterReceptionServiceServer,
			pb.RegisterLoginServiceServer,
		),
	)
}
