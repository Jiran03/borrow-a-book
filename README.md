
# Borrow a Book

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Jiran03/borrow-a-book)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Jiran03/borrow-a-book)


## Overview
Peminjaman buku

## Repository overview
Gambaran umum struktur direktori dan file:
```bash
├── README.md
├── go.mod
├── go.sum
├── .env
├── main.go
├───config
├───routes
├───auth
│   └───middlewares
├───book
│   ├───domain
│   ├───handler
│   │   └───api
│   ├───repository
│   │   └───mysql
│   └───service
├───borrow
│   ├───domain
│   ├───handler
│   │   └───api
│   ├───repository
│   │   └───mysql
│   └───service
├───helpers
│   ├───encrypt
│   ├───error
│   └───time
└───user
    ├───domain
    ├───handler
    │   └───api
    ├───repository
    │   └───mysql
    └───service
```

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
```