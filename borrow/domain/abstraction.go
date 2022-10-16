package domain

type Service interface {
	InsertData(domain Borrow) (borrowObj Borrow, err error)
	UpdateData(id string, domain Borrow) (borrowObj Borrow, err error)
	GetAllData() (borrowObj []Borrow, err error)
	GetByID(id string) (borrowObj Borrow, err error)
	DeleteData(id string) (err error)
}

type Repository interface {
	Create(domain Borrow) (borrowObj Borrow, err error)
	Update(domain Borrow) (borrowObj Borrow, err error)
	Get() (borrowObj []Borrow, err error)
	GetByID(id string) (domain Borrow, err error)
	Delete(id string) (err error)
}
