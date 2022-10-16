package service

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	bookDomain "github.com/Jiran03/borrow-a-book/book/domain"
	"github.com/Jiran03/borrow-a-book/borrow/domain"
	timeHelper "github.com/Jiran03/borrow-a-book/helpers/time"
	"github.com/google/uuid"
)

type borrowService struct {
	repository  domain.Repository
	bookService bookDomain.Service
	jwtAuth     authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (bs borrowService) UpdateData(id string, domain domain.Borrow) (borrowObj domain.Borrow, err error) {
	if borrowObj, err = bs.GetByID(id); err != nil {
		return borrowObj, err
	}

	domain.ID = borrowObj.ID
	domain.IDX = id
	domain.DeadlineAt = borrowObj.DeadlineAt
	domain.CreatedAt = borrowObj.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if borrowObj, err = bs.repository.Update(domain); err != nil {
		return borrowObj, err
	}

	return borrowObj, nil
}

// GetByID implements domain.Service
func (bs borrowService) GetByID(id string) (borrowObj domain.Borrow, err error) {
	if borrowObj, err = bs.repository.GetByID(id); err != nil {
		return borrowObj, err
	}

	return borrowObj, nil
}

func (bs borrowService) InsertData(domain domain.Borrow) (borrowObj domain.Borrow, err error) {
	domain.IDX = uuid.New().String()
	domain.DeadlineAt = timeHelper.DayToNano(domain.Duration)
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	if borrowObj, err = bs.repository.Create(domain); err != nil {
		return borrowObj, err
	}

	return borrowObj, nil
}

// GetAllData implements domain.Service
func (bs borrowService) GetAllData() (borrowObj []domain.Borrow, err error) {
	if borrowObj, err = bs.repository.Get(); err != nil {
		return borrowObj, err
	}

	return borrowObj, nil
}

// DeleteData implements domain.Service
func (bs borrowService) DeleteData(id string) (err error) {
	if errResp := bs.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewBorrowService(repo domain.Repository, bookService bookDomain.Service, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return borrowService{
		repository:  repo,
		bookService: bookService,
		jwtAuth:     jwtAuth,
	}
}