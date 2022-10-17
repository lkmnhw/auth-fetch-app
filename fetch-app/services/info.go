package services

import (
	"errors"
	"fetch-app/models"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func Info(c *http.Cookie) (*models.User, error) {
	// get value token
	claims, err := checkToken(c.Value)
	if err != nil {
		return nil, err
	}

	// assign value token to out
	user := extractToken(claims)
	return user, nil
}

func checkToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("Q4A9bPW8JcOqxpVzukO1"), nil
	})

	if err != nil {
		return claims, errors.New("Unauthorized")
	}

	if !token.Valid {
		return claims, errors.New("Unauthorized")
	}

	return claims, nil
}

func extractToken(claims jwt.MapClaims) *models.User {
	user := models.User{}

	for k, v := range claims {
		switch k {
		case "name":
			n, _ := v.(string)
			user.Name = n
		case "phone":
			p, _ := v.(string)
			user.Phone = p
		case "role":
			r, _ := v.(string)
			user.Role = r
		case "timestampp":
			t, _ := v.(float64)
			user.Timestampp = int64(t)
		}
	}

	return &user
}
