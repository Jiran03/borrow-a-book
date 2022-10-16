package borrow

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	repoMySQLB "github.com/Jiran03/borrow-a-book/book/repository/mysql"
	serviceB "github.com/Jiran03/borrow-a-book/book/service"
	handlerAPI "github.com/Jiran03/borrow-a-book/borrow/handler/api"
	repoMySQL "github.com/Jiran03/borrow-a-book/borrow/repository/mysql"
	service "github.com/Jiran03/borrow-a-book/borrow/service"
	"gorm.io/gorm"
)

func NewBorrowFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (borrowHandler handlerAPI.BorrowHandler) {
	bookRepo := repoMySQLB.NewBookRepository(db)
	bookService := serviceB.NewBookService(bookRepo, configJWT)
	borrowRepo := repoMySQL.NewBorrowRepository(db)
	borrowService := service.NewBorrowService(borrowRepo, bookService, configJWT)
	borrowHandler = handlerAPI.NewBorrowHandler(borrowService)
	return
}
