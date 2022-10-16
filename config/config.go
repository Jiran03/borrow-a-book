package config

import (
	"fmt"
	"os"

	bookRepo "github.com/Jiran03/borrow-a-book/book/repository/mysql"
	userRepo "github.com/Jiran03/borrow-a-book/user/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME: os.Getenv("DB_NAME"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASS"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
	}
}

func DBInit() (DB *gorm.DB) {

	//mysql
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)

	return

}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&userRepo.User{},
		&bookRepo.Book{},
	)
}
