package book

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	handlerAPI "github.com/Jiran03/borrow-a-book/book/handler/api"
	repoMySQL "github.com/Jiran03/borrow-a-book/book/repository/mysql"
	service "github.com/Jiran03/borrow-a-book/book/service"
	"gorm.io/gorm"
)

func NewBookFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (bookHandler handlerAPI.BookHandler) {
	bookRepo := repoMySQL.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo, configJWT)
	bookHandler = handlerAPI.NewBookHandler(bookService)
	return
}
