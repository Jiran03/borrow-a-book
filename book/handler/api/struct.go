package handlerAPI

import (
	"time"

	"github.com/Jiran03/borrow-a-book/book/domain"
	helperTime "github.com/Jiran03/borrow-a-book/helpers/time"
)

type RequestJSON struct {
	Title           string `json:"title" form:"title" validate:"required"`
	Writer          string `json:"writer" form:"writer" validate:"required"`
	Publisher       string `json:"publisher" form:"publisher" validate:"required"`
	PublicationYear int    `json:"publication_year" form:"publication_year" validate:"required"`
	BookQty         int    `json:"book_qty" form:"book_qty" validate:"required"`
}

func toDomain(req RequestJSON) domain.Book {
	return domain.Book{
		Title:           req.Title,
		Writer:          req.Writer,
		Publisher:       req.Publisher,
		PublicationYear: req.PublicationYear,
		BookQty:         req.BookQty,
	}
}

type ResponseJSON struct {
	ID              int       `json:"id"`
	IDX             string    `json:"idx"`
	Title           string    `json:"title" form:"title"`
	Writer          string    `json:"writer" form:"writer"`
	Publisher       string    `json:"publisher" form:"publisher"`
	PublicationYear int       `json:"publication_year" form:"publication_year"`
	BookQty         int       `json:"book_qty" form:"book_qty"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Book) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		ID:              domain.ID,
		IDX:             domain.IDX,
		Title:           domain.Title,
		Writer:          domain.Writer,
		Publisher:       domain.Publisher,
		PublicationYear: domain.PublicationYear,
		BookQty:         domain.BookQty,
		CreatedAt:       tmCreatedAt,
		UpdatedAt:       tmUpdatedAt,
	}
}
