package models

import "time"

type Shurl struct {
	Url    string        `json:"url"`
	Shurl  string        `json:"shurl"`
	Length int           `json:"length"`
	Expiry time.Duration `json:"expiry"`
}
