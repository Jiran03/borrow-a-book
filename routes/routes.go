package routes

import (
	"os"
	"strconv"

	"github.com/Jiran03/borrow-a-book/auth/middlewares"
	"github.com/Jiran03/borrow-a-book/book"
	"github.com/Jiran03/borrow-a-book/borrow"
	"github.com/Jiran03/borrow-a-book/config"
	"github.com/Jiran03/borrow-a-book/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	config.Init()
	db := config.DBInit()
	pool := config.NewPool()
	config.DBMigrate(db)
	expDuration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: expDuration,
	}

	user := user.NewUserFactory(db, configJWT)
	book := book.NewBookFactory(db, configJWT)
	borrow := borrow.NewBorrowFactory(db, pool, configJWT)
	e := echo.New()
	middlewares.LogMiddleware(e)
	v1 := e.Group("/v1")
	v1.POST("/register", user.Create)
	v1.POST("/login", user.Login)

	userG := v1.Group("/user")
	userG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	userG.GET("", user.GetAll)
	userG.GET("/:id", user.GetByID)
	userG.PATCH("/change-password/:id", user.ChangePassword)
	userG.PUT("/:id", user.Update, middlewares.UserValidation(user))
	userG.DELETE("/:id", user.Delete, middlewares.UserValidation(user))

	bookG := v1.Group("/book")
	bookG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	bookG.POST("", book.Create)
	bookG.GET("", book.GetAll)
	bookG.GET("/:id", book.GetByID)
	bookG.PUT("/:id", book.Update)
	bookG.DELETE("/:id", book.Delete)

	borrowG := v1.Group("/borrow")
	borrowG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	borrowG.POST("", borrow.Create)
	borrowG.GET("", borrow.GetAll)
	borrowG.GET("/:id", borrow.GetByID)
	borrowG.PUT("/:id", borrow.Update)
	borrowG.DELETE("/:id", borrow.Delete)

	return e
}
