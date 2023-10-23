package queries

import (
	"time"

	"github.com/axseem/shurl/database"
)

func GetShurl(id string) (string, error) {
	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, id).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func AddShurl(id string, url string, expiry time.Duration) error {
	r := database.CreateClient(0)
	defer r.Close()

	return r.Set(database.Ctx, id, url, expiry*3600*time.Second).Err()
}
