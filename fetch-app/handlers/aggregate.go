package handlers

import (
	"fetch-app/services"
	"log"
	"net/http"
	"time"
)

func (h Handler) Aggregate(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch list")
	strt := time.Now()
	out := Response{}

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
	if err != nil || user.Role != "admin" {
		out.Message = "Unauthorized"
		ResponseJSON(w, http.StatusUnauthorized, out)

		elpsd := time.Since(strt).Milliseconds()
		log.Println("\t benchmark: ", elpsd)
		return
	}

	agrgts, err := services.Aggregate()
	if err != nil {
		out.Message = err.Error()
		ResponseJSON(w, http.StatusInternalServerError, out)

		elpsd := time.Since(strt).Milliseconds()
		log.Println("\t benchmark: ", elpsd)
		return
	}

	out.Message = "success"
	out.Aggregates = agrgts
	ResponseJSON(w, http.StatusOK, out)

	elpsd := time.Since(strt).Milliseconds()
	log.Println("\t benchmark: ", elpsd)
}
