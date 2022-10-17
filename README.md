
# Borrow a Book

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Jiran03/borrow-a-book)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Jiran03/borrow-a-book)


## Overview
Peminjaman buku

## Repository overview
Gambaran umum struktur direktori dan file:
```bash
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   main.go
│   README.md
│
├───auth
│   └───middlewares
│           auth.go
│           log.go
│
├───book
│   │   factory.go
│   │
│   ├───domain
│   │       abstraction.go
│   │       core.go
│   │
│   ├───handler
│   │   └───api
│   │           controller.go
│   │           struct.go
│   │
│   ├───repository
│   │   └───mysql
│   │           mysql.go
│   │           record.go
│   │
│   └───service
│           usecase.go
│
├───borrow
│   │   factory.go
│   │
│   ├───domain
│   │       abstraction.go
│   │       core.go
│   │
│   ├───handler
│   │   └───api
│   │           controller.go
│   │           struct.go
│   │
│   ├───repository
│   │   ├───mysql
│   │   │       mysql.go
│   │   │       record.go
│   │   │
│   │   └───redis
│   │           record.go
│   │           redis.go
│   │
│   └───service
│           usecase.go
│
├───config
│       config.go
│
├───helpers
│   ├───encrypt
│   │       bcrypt.go
│   │
│   ├───error
│   │       error.go
│   │
│   └───time
│           parse.go
│           unix.go
│
├───routes
│       routes.go
│
└───user
    │   factory.go
    │
    ├───domain
    │       abstraction.go
    │       core.go
    │
    ├───handler
    │   └───api
    │           controller.go
    │           struct.go
    │
    ├───repository
    │   └───mysql
    │           mysql.go
    │           record.go
    │
    └───service
            usecase.go
```

### Software architecture & pattern
Aplikasi ini dibangun dan didesain menggunakan desain Clean Architecture dengan struktur folder seperti di atas. Clean architecture digunakan agar tiap layernya dapat berdiri sendiri sehingga tidak ada yang akan bentrok ketika ada perubahan code pada layer tententu. Struktur folder saya kelompokkan berdasarkan entitas agar mudah dalam memanajemen code serta melakukan testing. 

Agar tiap layer dapat berkomunikasi, saya menggunakan factory pattern. Pattern ini saya gunakan agar dapat menghindari pengaksesan kode secara langsung yang mungkin dapat membahayakan sistem.    


## .env file
Struktur env file:
```
//mysql
DB_USER="" //DB username
DB_PASS="" //DB password
DB_PORT=3306 //DB port
DB_HOST="" //DB host
DB_NAME="" //DB name

JWT_SECRET = "" //JWT secret key
JWT_EXPIRED = 1 //JWT expired

REDIHOST="" //redis host
REDIPORT=6379 //redis port
```