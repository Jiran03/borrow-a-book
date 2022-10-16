package service

import (
	"errors"

	authMiddleware "github.com/Jiran03/borrow-a-book/auth/middlewares"
	encryptHelper "github.com/Jiran03/borrow-a-book/helpers/encrypt"
	timeHelper "github.com/Jiran03/borrow-a-book/helpers/time"
	"github.com/Jiran03/borrow-a-book/user/domain"
	"github.com/google/uuid"
)

type userService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// ChangePassword implements domain.Service
func (us userService) ChangePassword(id string, domain domain.User) (userObj domain.User, err error) {
	if userObj, err = us.GetByID(id); err != nil {
		return userObj, err
	}

	if userObj.Password, err = encryptHelper.Hash(domain.Password); err != nil {
		return userObj, err
	}

	userObj.UpdatedAt = timeHelper.Timestamp()
	if userObj, err = us.repository.Update(userObj); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// UpdateData implements domain.Service
func (us userService) UpdateData(id string, domain domain.User) (userObj domain.User, err error) {
	if userObj, err = us.GetByID(id); err != nil {
		return userObj, err
	}

	if domain.Password, err = encryptHelper.Hash(domain.Password); err != nil {
		return userObj, err
	}

	domain.ID = userObj.ID
	domain.IDX = id
	domain.CreatedAt = userObj.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if userObj, err = us.repository.Update(domain); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByEmail implements domain.Service
func (us userService) GetByEmail(email string) (userObj domain.User, err error) {
	if userObj, err = us.repository.GetByEmail(email); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByID implements domain.Service
func (us userService) GetByID(id string) (userObj domain.User, err error) {
	if userObj, err = us.repository.GetByID(id); err != nil {
		return userObj, err
	}

	return userObj, nil
}

func (us userService) CreateToken(email, password string) (token string, userObj domain.User, err error) {
	if userObj, err = us.repository.GetByEmail(email); err != nil {
		return token, userObj, err
	}

	if !encryptHelper.ValidateHash(password, userObj.Password) {
		return token, userObj, errors.New("nama pengguna atau kata sandi salah")
	}

	id := userObj.IDX
	if token, err = us.jwtAuth.CreateToken(id, email); err != nil {
		return token, userObj, err
	}

	if userObj, err = us.GetByID(id); err != nil {
		return token, userObj, err
	}

	return token, userObj, nil
}

func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	email := domain.Email
	if _, errGetUser := us.repository.GetByEmail(email); errGetUser == nil {
		return userObj, errors.New("email telah terdaftar")
	}

	if domain.Password, err = encryptHelper.Hash(domain.Password); err != nil {
		return userObj, err
	}

	domain.IDX = uuid.New().String()
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	if userObj, err = us.repository.Create(domain); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetAllData implements domain.Service
func (us userService) GetAllData() (userObj []domain.User, err error) {
	if userObj, err = us.repository.Get(); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// DeleteData implements domain.Service
func (us userService) DeleteData(id int) (err error) {
	if errResp := us.repository.Delete(id); errResp != nil {
		return errResp
	}

	return nil
}

func NewUserService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return userService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
