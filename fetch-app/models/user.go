package models

import "time"

type User struct {
	Phone      string    `json:"phone" bson:"phone"`
	Name       string    `json:"name" bson:"name"`
	Role       string    `json:"role" bson:"role"`
	Password   *string   `json:"password,omitempty" bson:"password"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Timestampp int64     `json:"timestampp" bson:"timestampp"`
}
