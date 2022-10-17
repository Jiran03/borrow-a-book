package repoMySQL

import (
	"github.com/Jiran03/borrow-a-book/borrow/domain"
	"gorm.io/gorm"
)

type borrowRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (br borrowRepository) Update(domain domain.Borrow) (borrowObj domain.Borrow, err error) {
	var newRecord Borrow
	rec := fromDomain(domain)
	if err = br.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":          rec.ID,
		"id_x":        rec.IDX,
		"user_id":     rec.UserID,
		"book_id":     rec.BookID,
		"qty":         rec.Qty,
		"duration":    rec.Duration,
		"deadline_at": rec.DeadlineAt,
		"created_at":  rec.CreatedAt,
		"updated_at":  rec.UpdatedAt,
	}).Error; err != nil {
		return borrowObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (br borrowRepository) GetByID(id string) (domain domain.Borrow, err error) {
	var record Borrow
	err = br.DB.Where("id_x = ?", id).First(&record).Error
	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

// Get implements domain.Repository
func (br borrowRepository) Get() (borrowObj []domain.Borrow, err error) {
	var records []Borrow
	err = br.DB.Find(&records).Error
	if err != nil {
		return borrowObj, err
	}

	for _, value := range records {
		borrowObj = append(borrowObj, toDomain(value))
	}

	return borrowObj, nil
}

// Create implements domain.Repository
func (br borrowRepository) Create(domain domain.Borrow) (borrowObj domain.Borrow, err error) {
	// var recordDetail BorrowDetail
	record := fromDomain(domain)
	err = br.DB.Create(&record).Error
	if err != nil {
		return borrowObj, err
	}

	borrowObj = toDomain(record)
	return borrowObj, nil
}

// Delete implements domain.Repository
func (br borrowRepository) Delete(id string) (err error) {
	var record Borrow
	err = br.DB.Where("id_x = ?", id).Delete(&record).Error
	if err != nil {
		return err
	}

	return nil
}

func NewBorrowRepository(db *gorm.DB) domain.Repository {
	return borrowRepository{
		DB: db,
	}
}
