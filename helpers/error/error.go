package error

import (
	"errors"
	"net/http"
	"strings"
)

type QueryError struct {
	Query string
	Err   error
}

const (
	validation     = "validation"
	email          = "email"
	password       = "Password"
	server         = "Server"
	recordNotFound = "not found"
)

const (
	status400 = http.StatusBadRequest
	status404 = http.StatusNotFound
	status200 = http.StatusOK
	status500 = http.StatusInternalServerError
)

func ErrorMessage(err string) (code int, message error) {
	//validation user error
	if strings.Contains(err, validation) {
		if strings.Contains(err, password) {
			return status400, errors.New("password minimal terdiri dari 8 karakter")
		}
		if strings.Contains(err, email) {
			return status400, errors.New("email tidak sesuai format")
		}

		return status400, errors.New("isi semua field")
	}

	//server error
	if strings.Contains(err, recordNotFound) {
		return status404, errors.New("data tidak ditemukan")
	}
	if strings.Contains(err, server) {
		return status500, errors.New("internal server error")
	}

	return status200, errors.New("")
}
