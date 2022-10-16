package domain

type Service interface {
	InsertData(domain Book) (bookObj Book, err error)
	UpdateData(id string, domain Book) (bookObj Book, err error)
	GetAllData() (bookObj []Book, err error)
	GetByID(id string) (bookObj Book, err error)
	DeleteData(id string) (err error)
}

type Repository interface {
	Create(domain Book) (bookObj Book, err error)
	Update(domain Book) (bookObj Book, err error)
	Get() (bookObj []Book, err error)
	GetByID(id string) (domain Book, err error)
	Delete(id string) (err error)
}
