package handlers

import (
	"auth-app/models"
	"auth-app/services"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("login")
	strt := time.Now()

	// parse input to model user
	in := models.User{}
	out := Response{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		out.Message = err.Error()
		ResponseJSON(w, http.StatusBadRequest, out)

		elpsd := time.Since(strt).Milliseconds()
		log.Println("\t benchmark: ", elpsd)
		return
	}
	defer r.Body.Close()

	// service login to verify login and get token
	token, err := services.Login(h.Context, h.Database, in)
	if err != nil {
		out.Message = err.Error()
		ResponseJSON(w, http.StatusBadRequest, out)

		elpsd := time.Since(strt).Milliseconds()
		log.Println("\t benchmark: ", elpsd)
		return
	}

	// set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	out.Message = "success"
	ResponseJSON(w, http.StatusOK, out)

	elpsd := time.Since(strt).Milliseconds()
	log.Println("\t benchmark: ", elpsd)
}
