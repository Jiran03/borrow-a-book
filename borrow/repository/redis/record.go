package repoRedis

import "github.com/Jiran03/borrow-a-book/borrow/domain"

type Borrow struct {
	ID         int    `redis:"id"`
	IDX        string `redis:"id_x"`
	UserID     int    `redis:"user_id"`
	BookID     int    `redis:"book_id"`
	Qty        int    `redis:"qty"`
	Duration   int    `redis:"duration"`
	DeadlineAt string `redis:"deadline_at"`
	CreatedAt  string `redis:"created_at"`
	UpdatedAt  string `redis:"updated_at"`
}

func toDomain(rec Borrow) domain.Borrow {
	return domain.Borrow{
		ID:         rec.ID,
		IDX:        rec.IDX,
		UserID:     rec.UserID,
		BookID:     rec.BookID,
		Qty:        rec.Qty,
		Duration:   rec.Duration,
		DeadlineAt: rec.DeadlineAt,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}
