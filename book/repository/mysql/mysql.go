package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/book/domain"
	"gorm.io/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (ur bookRepository) Update(domain domain.Book) (bookObj domain.Book, err error) {
	var newRecord Book
	rec := fromDomain(domain)
	if err = ur.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":               rec.ID,
		"id_x":             rec.IDX,
		"title":            rec.Title,
		"writer":           rec.Writer,
		"publisher":        rec.Publisher,
		"publication_year": rec.PublicationYear,
		"book_qty":         rec.BookQty,
		"created_at":       rec.CreatedAt,
		"updated_at":       rec.UpdatedAt,
	}).Error; err != nil {
		return bookObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (ur bookRepository) GetByID(id string) (domain domain.Book, err error) {
	var record Book
	err = ur.DB.Where("id_x = ?", id).First(&record).Error
	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (ur bookRepository) Get() (bookObj []domain.Book, err error) {
	var records []Book
	err = ur.DB.Find(&records).Error
	if err != nil {
		return bookObj, err
	}

	for _, value := range records {
		bookObj = append(bookObj, toDomain(value))
	}

	return bookObj, nil
}

// Create implements domain.Repository
func (ur bookRepository) Create(domain domain.Book) (bookObj domain.Book, err error) {
	// var recordDetail BookDetail
	record := fromDomain(domain)
	err = ur.DB.Create(&record).Error
	if err != nil {
		return bookObj, err
	}

	bookObj = toDomain(record)
	return bookObj, nil
}

// Delete implements domain.Repository
func (ur bookRepository) Delete(id string) (err error) {
	var record Book
	err = ur.DB.Where("id_x = ?", id).Delete(&record).Error
	if err != nil {
		return err
	}

	return nil
}

func NewBookRepository(db *gorm.DB) domain.Repository {
	return bookRepository{
		DB: db,
	}
}
