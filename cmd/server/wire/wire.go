//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/internal/handler"
	categoryHandler "github.com/ch3nnn/webstack-go/internal/handler/category"
	dashboardHandler "github.com/ch3nnn/webstack-go/internal/handler/dashboard"
	indexHandler "github.com/ch3nnn/webstack-go/internal/handler/index"
	siteHandler "github.com/ch3nnn/webstack-go/internal/handler/site"
	userHandler "github.com/ch3nnn/webstack-go/internal/handler/user"
	"github.com/ch3nnn/webstack-go/internal/server"
	"github.com/ch3nnn/webstack-go/internal/service"
	categoryService "github.com/ch3nnn/webstack-go/internal/service/category"
	indexService "github.com/ch3nnn/webstack-go/internal/service/index"
	siteService "github.com/ch3nnn/webstack-go/internal/service/site"
	userService "github.com/ch3nnn/webstack-go/internal/service/user"
	"github.com/ch3nnn/webstack-go/pkg/app"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
	"github.com/ch3nnn/webstack-go/pkg/server/http"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	userHandler.NewHandler,
	indexHandler.NewHandler,
	siteHandler.NewHandler,
	categoryHandler.NewHandler,
	dashboardHandler.NewHandler,
)

var serviceSet = wire.NewSet(
	service.NewService,
	userService.NewService,
	indexService.NewService,
	siteService.NewService,
	categoryService.NewService,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

// build App
func newApp(httpServer *http.Server) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("webstack-go"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		serviceSet,
		handlerSet,
		repositorySet,
		jwt.NewJwt,
		http.NewGinDefaultServer,
		newApp,
	))
}