//go:build wireinject
// +build wireinject

package server

import (
	config "UserService/config"
	handlers "UserService/internal/handlers"
	repository "UserService/internal/repository"
	services "UserService/internal/services"

	"github.com/google/wire"
)

// Wire khởi tạo tất cả các phụ thuộc và trả về UserHandler
func InitializeUserHandler() *handlers.UserHandler {
	wire.Build(
		config.InitDB,
		repository.NewUserRepository,
		services.NewUserService,
		handlers.NewUserHandler,
	)
	return nil
}
