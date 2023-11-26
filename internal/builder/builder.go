package builder

import (
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/router"
	"gorm.io/gorm"
)

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// userRepository := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := handler.NewUserHandler(userService)
	// return router.PrivateRoutes(userHandler)
	return nil
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// 	userRepository := repository.NewUserRepository(db)
	// 	loginService := service.NewLoginService(userRepository)
	// 	tokenService := service.NewTokenService(cfg)
	// 	authHandler := handler.NewAuthHandler(loginService, tokenService)
	// 	return router.PublicRoutes(authHandler)
	return nil
}
