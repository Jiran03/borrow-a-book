package repoRedis

import (
	"encoding/json"
	"time"

	"github.com/Jiran03/borrow-a-book/borrow/domain"
	"github.com/gomodule/redigo/redis"
)

type BorrowCache struct {
	pool *redis.Pool
}

func (bc BorrowCache) Delete() (err error) {
	client := bc.pool.Get()
	defer client.Close()
	_, err = client.Do("FLUSHALL")
	if err != nil {
		return err
	}

	return nil
}

// Create
func (bc BorrowCache) Create(domain interface{}) (err error) {
	client := bc.pool.Get()
	defer client.Close()
	// client.Do("FLUSHALL")
	rec, _ := json.Marshal(domain)
	_, err = redis.DoWithTimeout(client, 1*time.Minute, "APPEND", "borrow:1", rec)
	if err != nil {
		return err
	}

	return nil
}

// Get
func (bc BorrowCache) Get() (borrowObj []domain.Borrow, err error) {
	client := bc.pool.Get()
	defer client.Close()

	var rec = []Borrow{}
	str, _ := redis.Bytes(client.Do("GET", "borrow:1"))
	json.Unmarshal(str, &rec)
	if err != nil {
		return borrowObj, err
	}

	for _, v := range rec {
		borrowObj = append(borrowObj, toDomain(v))
	}

	return borrowObj, nil
}

func NewBorrowCache(pool *redis.Pool) BorrowCache {
	return BorrowCache{
		pool: pool,
	}
}
