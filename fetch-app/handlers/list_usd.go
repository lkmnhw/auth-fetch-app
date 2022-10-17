package handlers

import (
	"fetch-app/services"
	"log"
	"net/http"
	"time"
)

func (h Handler) ListUSD(w http.ResponseWriter, r *http.Request) {
	log.Println("fetch list")
	strt := time.Now()
	out := Response{}

	cmdts, err := services.FetchListUSD(h.Cache)
	if err != nil {
		out.Message = err.Error()
		ResponseJSON(w, http.StatusInternalServerError, out)
		return
	}

	out.Message = "success"
	out.Commodities = cmdts
	ResponseJSON(w, http.StatusOK, out)

	elpsd := time.Since(strt).Milliseconds()
	log.Println("\t benchmark: ", elpsd)
}
