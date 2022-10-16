package user

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	handlerAPI "github.com/Jiran03/borrow-a-book/user/handler/api"
	repoMySQL "github.com/Jiran03/borrow-a-book/user/repository/mysql"
	service "github.com/Jiran03/borrow-a-book/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (userHandler handlerAPI.UserHandler) {
	userRepo := repoMySQL.NewUserRepository(db)
	userService := service.NewUserService(userRepo, configJWT)
	userHandler = handlerAPI.NewUserHandler(userService)
	return
}
