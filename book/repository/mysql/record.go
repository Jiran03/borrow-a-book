package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/book/domain"
	repoMySQLB "github.com/Jiran03/borrow-a-book/borrow/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              int    `gorm:"index"`
	IDX             string `gorm:"index:id_x,unique"`
	Title           string
	Writer          string
	Publisher       string
	PublicationYear int
	BookQty         int
	CreatedAt       string
	UpdatedAt       string
	Borrows         []repoMySQLB.Borrow `gorm:"foreignKey:BookID"`
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
