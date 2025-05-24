package entity

import "time"

type Order struct {
	ID        int
	UserID    int64
	Service   string
	Date      string
	Contact   string
	CreatedAt time.Time
}
