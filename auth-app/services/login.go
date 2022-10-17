package services

import (
	"context"
	"errors"
	"log"
	"time"

	"auth-app/databases"
	"auth-app/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaim struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Timestampp int64  `json:"timestampp"`
	jwt.RegisteredClaims
}

func Login(ctx context.Context, db *databases.Database, in models.User) (string, error) {
	// validate input
	if valid := userValidation(in); valid != "" {
		return "", errors.New(valid)
	}

	// check if user exist
	user, err := models.GetUser(ctx, db, in.Phone)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New(`wrong user or password`)
	}

	// parse user record and check if password nil
	if user.Password == nil {
		log.Printf("user %s password is not set", user.Name)
		return "", errors.New(`wrong user or password`)
	}

	// cek if password valid
	if err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(*in.Password),
	); err != nil {
		return "", errors.New(`wrong user or password`)
	}

	// set jwt
	token, err := createJwt(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createJwt(user models.User) (string, error) {
	claims := &JWTClaim{
		Name:       user.Name,
		Phone:      user.Phone,
		Role:       user.Role,
		Timestampp: user.Timestampp,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-app",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("Q4A9bPW8JcOqxpVzukO1"))
	if err != nil {
		return "", err
	}
	return token, nil
}
