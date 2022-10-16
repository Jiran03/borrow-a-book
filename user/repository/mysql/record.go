package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	IDX       string
	Name      string
	Email     string
	Gender    string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec User) domain.User {
	return domain.User{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Gender:    rec.Gender,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		ID:        rec.ID,
		IDX:       rec.IDX,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Gender:    rec.Gender,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
