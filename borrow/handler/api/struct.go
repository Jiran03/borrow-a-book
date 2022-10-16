package handlerAPI

import (
	"time"

	"github.com/Jiran03/borrow-a-book/borrow/domain"
	helperTime "github.com/Jiran03/borrow-a-book/helpers/time"
)

type RequestJSON struct {
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	BookID   int `json:"book_id" form:"book_id" validate:"required"`
	Qty      int `json:"qty" form:"qty" validate:"required"`
	Duration int `json:"duration" form:"duration" validate:"required"`
}

func toDomain(req RequestJSON) domain.Borrow {
	return domain.Borrow{
		UserID:   req.UserID,
		BookID:   req.BookID,
		Qty:      req.Qty,
		Duration: req.Duration,
	}
}

type ResponseJSON struct {
	ID         int       `json:"id"`
	IDX        string    `json:"idx"`
	UserID     int       `json:"user_id" form:"user_id" validate:"required"`
	BookID     int       `json:"book_id" form:"book_id" validate:"required"`
	Qty        int       `json:"qty" form:"qty" validate:"required"`
	DeadlineAt time.Time `json:"deadline_at" form:"deadline_at" validate:"required"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Borrow) ResponseJSON {
	//parse unix timestamp to time.Time
	tmDeadlineAt := helperTime.NanoToTime(domain.DeadlineAt)
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:         domain.ID,
		IDX:        domain.IDX,
		UserID:     domain.UserID,
		BookID:     domain.BookID,
		Qty:        domain.Qty,
		DeadlineAt: tmDeadlineAt,
		CreatedAt:  tmCreatedAt,
		UpdatedAt:  tmUpdatedAt,
	}
}
