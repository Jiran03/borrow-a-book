
# Borrow a Book

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Jiran03/borrow-a-book)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Jiran03/borrow-a-book)



## Overview
Aplikasi peminjaman buku. Aplikasi masih berjalan di lokal dengan base url `localhost:8080/v1`. Gunakan [Postman](https://www.postman.com/downloads/) untuk pengujian aplikasi.



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



## Installation


### Prepare

| **Tools**         | **Versi**  |
|-------------------|-----------|
| Go                | [1.18](https://go.dev/dl/go1.18.windows-amd64.msi)      |
| Echo                | [v4](https://echo.labstack.com/guide/#installation)      |
| Gorm          | [v2.0.0](https://gorm.io/docs/v2_release_note.html#Install)      |
| Redigo          | [v1.8.9](https://github.com/gomodule/redigo)      |
| MySQL          | [latest](https://dev.mysql.com/downloads/mysql/)      |
| Redis          | [7.0](https://redis.io/download/)      |
| Postman          | [v10](https://www.postman.com/downloads/)      |


### .env file

Buat sebuah file bernama `.env` kemudian sesuaikan seperti berikut:
```
//mysql
DB_USER = "" //DB username
DB_PASS = "" //DB password
DB_PORT = 3306 //DB port
DB_HOST = "" //DB host
DB_NAME = "" //DB name

JWT_SECRET = "" //JWT secret key
JWT_EXPIRED = 1 //JWT expired

REDIHOST = "" //redis host
REDIPORT = 6379 //redis port
```
Sesuaikan value-nya dengan database yang sudah Anda buat.



## Run & Testing app


Gunakan perintah  `go run main.go` untuk menjalankan aplikasi. Untuk melakukan testing silakan gunakan postman kemudian import file berikut: [e2e test file]("github.com/Jiran03/borrow-a-book/BorrowABook.postman_collection.json")