package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/borrow/domain"
	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	ID         int    `gorm:"index"`
	IDX        string `gorm:"index:id_x,unique"`
	UserID     int
	BookID     int
	Qty        int
	Duration   int
	DeadlineAt string
	CreatedAt  string
	UpdatedAt  string
}

func toDomain(rec Borrow) domain.Borrow {
	return domain.Borrow{
		ID:         rec.ID,
		IDX:        rec.IDX,
		UserID:     rec.UserID,
		BookID:     rec.BookID,
		Qty:        rec.Qty,
		Duration:   rec.Duration,
		DeadlineAt: rec.DeadlineAt,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Borrow) Borrow {
	return Borrow{
		ID:         rec.ID,
		IDX:        rec.IDX,
		UserID:     rec.UserID,
		BookID:     rec.BookID,
		Qty:        rec.Qty,
		Duration:   rec.Duration,
		DeadlineAt: rec.DeadlineAt,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}
