package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/borrow-a-book/user/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (uh UserHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	err := uh.validation.Struct(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err = uh.service.InsertData(toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}

func (uh UserHandler) Login(ctx echo.Context) error {
	var req RequestLoginJSON
	ctx.Bind(&req)
	err := uh.validation.Struct(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	email := req.Email
	password := req.Password
	token, userRes, err := uh.service.CreateToken(email, password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"token":   token,
		"data":    userObj,
	})
}

func (uh UserHandler) GetAll(ctx echo.Context) error {
	userRes, err := uh.service.GetAllData()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := []ResponseJSON{}
	for _, value := range userRes {
		userObj = append(userObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	userRes, err := uh.service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	userRes, err := uh.service.UpdateData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetByEmail(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	userRes, err := uh.service.GetByEmail(email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := uh.service.DeleteData(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}

func (uh UserHandler) ChangePassword(ctx echo.Context) error {
	var req RequestPasswordJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	err := uh.validation.Struct(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	userRes, err := uh.service.ChangePassword(id, pwdToDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetValidEmail(email string) (id string, err error) {
	userObj, err := uh.service.GetByEmail(email)
	if err != nil {
		return userObj.IDX, err
	}

	return userObj.IDX, nil
}
