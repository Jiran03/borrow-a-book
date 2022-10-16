package domain

type Service interface {
	CreateToken(email, password string) (token string, userObj User, err error)
	InsertData(domain User) (userObj User, err error)
	UpdateData(id string, domain User) (userObj User, err error)
	GetAllData() (userObj []User, err error)
	GetByID(id string) (userObj User, err error)
	GetByEmail(email string) (userObj User, err error)
	DeleteData(id string) (err error)
	ChangePassword(id string, domain User) (userObj User, err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	Update(domain User) (userObj User, err error)
	Get() (userObj []User, err error)
	GetByID(id string) (domain User, err error)
	GetByEmail(email string) (userObj User, err error)
	Delete(id string) (err error)
}
