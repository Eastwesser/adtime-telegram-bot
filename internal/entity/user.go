package entity

import "time"

type User struct {
	ID        int64
	Name      string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Consent   bool
	CreatedAt time.Time
}
