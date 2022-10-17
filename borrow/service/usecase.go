package service

import (
	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	bookDomain "github.com/Jiran03/borrow-a-book/book/domain"
	"github.com/Jiran03/borrow-a-book/borrow/domain"
	repoRedis "github.com/Jiran03/borrow-a-book/borrow/repository/redis"
	timeHelper "github.com/Jiran03/borrow-a-book/helpers/time"
	"github.com/google/uuid"
)

type borrowService struct {
	repository  domain.Repository
	cache       repoRedis.BorrowCache
	bookService bookDomain.Service
	jwtAuth     authMiddleware.ConfigJWT
}

// UpdateData implements domain.Service
func (bs borrowService) UpdateData(id string, domain domain.Borrow) (borrowObj domain.Borrow, err error) {
	if err = bs.cache.Delete(); err != nil {
		return borrowObj, err
	}

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
	if err = bs.cache.Delete(); err != nil {
		return borrowObj, err
	}

	bookObj, err := bs.bookService.GetByID(domain.BookID)
	if err != nil {
		return borrowObj, err
	}

	if bookObj.BookQty < domain.Qty {
		return borrowObj, err
	}

	bookObj.BookQty -= domain.Qty
	if _, err = bs.bookService.UpdateData(bookObj.IDX, bookObj); err != nil {
		return borrowObj, err
	}

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
	borrowObj, _ = bs.cache.Get()
	if len(borrowObj) == 0 {
		borrowObj, err = bs.repository.Get()
		if err != nil {
			return borrowObj, err
		}

		err = bs.cache.Create(borrowObj)
		if err != nil {
			return borrowObj, err
		}

		borrowObj, err = bs.cache.Get()
		if err != nil {
			return borrowObj, err
		}
	}

	return borrowObj, nil
}

// DeleteData implements domain.Service
func (bs borrowService) DeleteData(id string) (err error) {
	if err = bs.cache.Delete(); err != nil {
		return err
	}

	if errResp := bs.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewBorrowService(repo domain.Repository, pool repoRedis.BorrowCache, bookService bookDomain.Service, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return borrowService{
		repository:  repo,
		cache:       pool,
		bookService: bookService,
		jwtAuth:     jwtAuth,
	}
}
