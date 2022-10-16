package handlerAPI

import (
	"fmt"
	"net/http"

	"github.com/Jiran03/borrow-a-book/borrow/domain"
	errHelper "github.com/Jiran03/borrow-a-book/helpers/error"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BorrowHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewBorrowHandler(service domain.Service) BorrowHandler {
	return BorrowHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (bh BorrowHandler) Create(ctx echo.Context) error {
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

func (bh BorrowHandler) GetAll(ctx echo.Context) error {
	borrowRes, err := bh.service.GetAllData()
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	borrowObj := []ResponseJSON{}
	for _, value := range borrowRes {
		borrowObj = append(borrowObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    borrowObj,
	})
}

func (bh BorrowHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	borrowRes, err := bh.service.GetByID(id)
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	borrowObj := fromDomain(borrowRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    borrowObj,
	})
}

func (bh BorrowHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	borrowRes, err := bh.service.UpdateData(id, toDomain(req))
	if err != nil {
		errCode, errMessage := errHelper.ErrorMessage(err.Error())
		return ctx.JSON(errCode, map[string]interface{}{
			"message": errMessage.Error(),
			"rescode": errCode,
		})
	}

	borrowObj := fromDomain(borrowRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    borrowObj,
	})
}

func (bh BorrowHandler) Delete(ctx echo.Context) error {
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
