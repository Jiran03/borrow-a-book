package handlerAPI

import (
	"fmt"
	"net/http"

	"github.com/Jiran03/borrow-a-book/book/domain"
	errHelper "github.com/Jiran03/borrow-a-book/helpers/error"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewBookHandler(service domain.Service) BookHandler {
	return BookHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (bh BookHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	err := bh.validation.Struct(req)
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	_, err = bh.service.InsertData(toDomain(req))
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}

func (bh BookHandler) GetAll(ctx echo.Context) error {
	bookRes, err := bh.service.GetAllData()
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	bookObj := []ResponseJSON{}
	for _, value := range bookRes {
		bookObj = append(bookObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	bookRes, err := bh.service.GetByID(id)
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	bookObj := fromDomain(bookRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	bookRes, err := bh.service.UpdateData(id, toDomain(req))
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	bookObj := fromDomain(bookRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := bh.service.DeleteData(id)
	if err != nil {
		fmt.Println(err)
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errCode,
			"rescode": errMessage.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}
