package services

import (
	"context"
	"errors"
	"math/rand"

	"auth-app/databases"
	"auth-app/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, db *databases.Database, in models.User) (*models.User, error) {
	// validate input
	if errValid := userValidation(in); errValid != "" {
		return nil, errors.New(errValid)
	}

	// get existing user
	user, _ := models.GetUser(ctx, db, in.Phone)
	if user != nil {
		return nil, errors.New(`user already exist`)
	}

	// generate password random string(4)
	if in.Password == nil {
		pswd := randString(4)
		in.Password = &pswd
	}
	oldPswd := *in.Password

	// hash password
	in.Password = hashPassword(*in.Password)

	// insert user to database
	newuser, err := models.InsertUser(ctx, db, in)
	if err != nil {
		return nil, err
	}

	newuser.Password = &oldPswd
	return &newuser, nil
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func hashPassword(pswd string) *string {
	hashPswd, _ := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	newPswd := string(hashPswd)
	return &newPswd
}
