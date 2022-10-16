package domain

type Borrow struct {
	ID         int
	IDX        string
	UserID     int
	BookID     int
	Qty        int
	Duration   int
	DeadlineAt string
	CreatedAt  string
	UpdatedAt  string
}
