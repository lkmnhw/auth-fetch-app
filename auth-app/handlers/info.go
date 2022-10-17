package handlers

import (
	"auth-app/services"
	"log"
	"net/http"
	"time"
)

func (h Handler) Info(w http.ResponseWriter, r *http.Request) {
	log.Println("info")
	strt := time.Now()

	out := Response{}
	// get jwt from cookie
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			out.Message = "Unauthorized"
			ResponseJSON(w, http.StatusUnauthorized, out)

			elpsd := time.Since(strt).Milliseconds()
			log.Println("\t benchmark: ", elpsd)
			return
		}
	}

	// service info get valid user from token
	user, err := services.Info(c)
	if err != nil {
		out.Message = "Unauthorized"
		ResponseJSON(w, http.StatusUnauthorized, out)

		elpsd := time.Since(strt).Milliseconds()
		log.Println("\t benchmark: ", elpsd)
		return
	}

	if user != nil {
		out.User = user
	}
	out.Message = "success"
	ResponseJSON(w, http.StatusOK, out)

	elpsd := time.Since(strt).Milliseconds()
	log.Println("\t benchmark: ", elpsd)
}
