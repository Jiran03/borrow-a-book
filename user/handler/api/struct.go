package handlerAPI

import (
	"time"

	helperTime "github.com/Jiran03/borrow-a-book/helpers/time"
	"github.com/Jiran03/borrow-a-book/user/domain"
)

type RequestJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
}

type RequestLoginJSON struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestPasswordJSON struct {
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Gender:   req.Gender,
	}
}

func pwdToDomain(req RequestPasswordJSON) domain.User {
	return domain.User{
		Password: req.Password,
	}
}

type ResponseJSON struct {
	ID        int       `json:"id"`
	IDX       string    `json:"idx"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Gender    string    `json:"gender" form:"gender"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:        domain.ID,
		IDX:       domain.IDX,
		Name:      domain.Name,
		Email:     domain.Email,
		Gender:    domain.Gender,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
