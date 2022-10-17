package config

import (
	"fmt"
	"log"
	"os"

	bookRepo "github.com/Jiran03/borrow-a-book/book/repository/mysql"
	borrowRepo "github.com/Jiran03/borrow-a-book/borrow/repository/mysql"
	userRepo "github.com/Jiran03/borrow-a-book/user/repository/mysql"
	"github.com/gomodule/redigo/redis"
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

type RedistConfig struct {
	REDIHOST string
	REDIPORT string
}

var Conf Config
var RediConf RedistConfig

func Init() {
	Conf = Config{
		DBNAME: os.Getenv("DB_NAME"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASS"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
	}

	RediConf = RedistConfig{
		REDIHOST: os.Getenv("REDIHOST"),
		REDIPORT: os.Getenv("REDIPORT"),
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
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

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			//redis
			rediAdd := fmt.Sprintf("%s:%s", RediConf.REDIHOST, RediConf.REDIPORT)
			c, err := redis.Dial("tcp", rediAdd)
			checkError(err)
			return c, err
		},
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&userRepo.User{},
		&bookRepo.Book{},
		&borrowRepo.Borrow{},
	)
}
