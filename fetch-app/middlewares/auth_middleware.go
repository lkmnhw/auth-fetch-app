package middlewares

import (
	"fetch-app/handlers"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// AuthRequired check if user already login (has token)
func AuthRequired(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out := handlers.Response{}

		// get jwt from cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				out.Message = "Unauthorized"
				handlers.ResponseJSON(w, http.StatusUnauthorized, out)
				return
			}
		}

		// get value token
		tokenString := c.Value
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("Q4A9bPW8JcOqxpVzukO1"), nil
		})

		if err != nil {
			out.Message = "Unauthorized"
			handlers.ResponseJSON(w, http.StatusUnauthorized, out)
			return
		}

		if !token.Valid {
			out.Message = "Unauthorized"
			handlers.ResponseJSON(w, http.StatusUnauthorized, out)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
