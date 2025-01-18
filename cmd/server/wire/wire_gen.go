// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/internal/handler"
	category2 "github.com/ch3nnn/webstack-go/internal/handler/category"
	config2 "github.com/ch3nnn/webstack-go/internal/handler/config"
	"github.com/ch3nnn/webstack-go/internal/handler/dashboard"
	index3 "github.com/ch3nnn/webstack-go/internal/handler/index"
	site2 "github.com/ch3nnn/webstack-go/internal/handler/site"
	user2 "github.com/ch3nnn/webstack-go/internal/handler/user"
	"github.com/ch3nnn/webstack-go/internal/server"
	"github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/internal/service/category"
	"github.com/ch3nnn/webstack-go/internal/service/config"
	index2 "github.com/ch3nnn/webstack-go/internal/service/index"
	"github.com/ch3nnn/webstack-go/internal/service/site"
	"github.com/ch3nnn/webstack-go/internal/service/user"
	"github.com/ch3nnn/webstack-go/pkg/app"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
	"github.com/ch3nnn/webstack-go/pkg/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	engine := http.NewGinDefaultServer()
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	indexHandler := index.NewHandler(handlerHandler)
	db := repository.NewDB(viperViper, logger)
	repositoryRepository := repository.NewRepository(logger, db)
	serviceService := service.NewService(engine, logger, jwtJWT, repositoryRepository)
	indexService := index2.NewService(serviceService)
	handler2 := index3.NewHandler(handlerHandler, indexService)
	userService := user.NewService(serviceService)
	userHandler := user2.NewHandler(handlerHandler, userService)
	siteService := site.NewService(serviceService)
	siteHandler := site2.NewHandler(handlerHandler, siteService)
	categoryService := category.NewService(serviceService)
	categoryHandler := category2.NewHandler(handlerHandler, categoryService)
	configService := config.NewService(serviceService)
	configHandler := config2.NewHandler(handlerHandler, configService)
	httpServer := server.NewHTTPServer(engine, logger, viperViper, jwtJWT, indexHandler, handler2, userHandler, siteHandler, categoryHandler, configHandler)
	appApp := newApp(httpServer)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository)

var handlerSet = wire.NewSet(handler.NewHandler, user2.NewHandler, index3.NewHandler, site2.NewHandler, category2.NewHandler, index.NewHandler, config2.NewHandler)

var serviceSet = wire.NewSet(service.NewService, user.NewService, index2.NewService, site.NewService, category.NewService, config.NewService)

var serverSet = wire.NewSet(server.NewHTTPServer)

// build App
func newApp(httpServer *http.Server) *app.App {
	return app.NewApp(app.WithServer(httpServer), app.WithName("webstack-go"))
}
