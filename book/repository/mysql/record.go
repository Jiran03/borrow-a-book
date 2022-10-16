package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/book/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              int
	IDX             string
	Title           string
	Writer          string
	Publisher       string
	PublicationYear int
	BookQty         int
	CreatedAt       string
	UpdatedAt       string
}

func toDomain(rec Book) domain.Book {
	return domain.Book{
		ID:              rec.ID,
		IDX:             rec.IDX,
		Title:           rec.Title,
		Writer:          rec.Writer,
		Publisher:       rec.Publisher,
		PublicationYear: rec.PublicationYear,
		BookQty:         rec.BookQty,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Book) Book {
	return Book{
		ID:              rec.ID,
		IDX:             rec.IDX,
		Title:           rec.Title,
		Writer:          rec.Writer,
		Publisher:       rec.Publisher,
		PublicationYear: rec.PublicationYear,
		BookQty:         rec.BookQty,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}
