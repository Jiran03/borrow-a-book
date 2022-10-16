package service

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	"github.com/Jiran03/borrow-a-book/book/domain"
	timeHelper "github.com/Jiran03/borrow-a-book/helpers/time"
	"github.com/google/uuid"
)

type bookService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (bs bookService) UpdateData(id string, domain domain.Book) (bookObj domain.Book, err error) {
	if bookObj, err = bs.GetByID(id); err != nil {
		return bookObj, err
	}

	domain.ID = bookObj.ID
	domain.IDX = id
	domain.CreatedAt = bookObj.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if bookObj, err = bs.repository.Update(domain); err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// GetByID implements domain.Service
func (bs bookService) GetByID(id string) (bookObj domain.Book, err error) {
	if bookObj, err = bs.repository.GetByID(id); err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

func (bs bookService) InsertData(domain domain.Book) (bookObj domain.Book, err error) {
	domain.IDX = uuid.New().String()
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	if bookObj, err = bs.repository.Create(domain); err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// GetAllData implements domain.Service
func (bs bookService) GetAllData() (bookObj []domain.Book, err error) {
	if bookObj, err = bs.repository.Get(); err != nil {
		return bookObj, err
	}

	return bookObj, nil
}

// DeleteData implements domain.Service
func (bs bookService) DeleteData(id string) (err error) {
	if errResp := bs.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewBookService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return bookService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
