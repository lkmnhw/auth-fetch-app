package models

import (
	"auth-app/databases"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Phone      string    `json:"phone" bson:"phone"`
	Name       string    `json:"name" bson:"name"`
	Role       string    `json:"role" bson:"role"`
	Password   *string   `json:"password,omitempty" bson:"password"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Timestampp int64     `json:"timestampp" bson:"timestampp"`
}

func GetUser(ctx context.Context, db *databases.Database, phone string) (*User, error) {
	r, err := db.FindOne(ctx, db.User, filterGetUser(phone))
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	user := User{}
	err = r.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func filterGetUser(phone string) bson.M {
	return bson.M{
		"phone": phone,
	}
}

func InsertUser(ctx context.Context, db *databases.Database, user User) (User, error) {
	user.CreatedAt = time.Now()
	user.Timestampp = user.CreatedAt.Unix()
	_, err := db.InsertOne(ctx, db.User, user)
	return user, err
}
