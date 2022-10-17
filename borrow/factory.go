package borrow

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	repoMySQLB "github.com/Jiran03/borrow-a-book/book/repository/mysql"
	serviceB "github.com/Jiran03/borrow-a-book/book/service"
	handlerAPI "github.com/Jiran03/borrow-a-book/borrow/handler/api"
	repoMySQL "github.com/Jiran03/borrow-a-book/borrow/repository/mysql"
	repoRedis "github.com/Jiran03/borrow-a-book/borrow/repository/redis"
	service "github.com/Jiran03/borrow-a-book/borrow/service"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

func NewBorrowFactory(db *gorm.DB, pool *redis.Pool, configJWT authMiddleware.ConfigJWT) (borrowHandler handlerAPI.BorrowHandler) {
	bookRepo := repoMySQLB.NewBookRepository(db)
	bookService := serviceB.NewBookService(bookRepo, configJWT)
	borrowRepo := repoMySQL.NewBorrowRepository(db)
	borrowCache := repoRedis.NewBorrowCache(pool)
	borrowService := service.NewBorrowService(borrowRepo, borrowCache, bookService, configJWT)
	borrowHandler = handlerAPI.NewBorrowHandler(borrowService)
	return
}
