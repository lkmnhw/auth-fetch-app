package servers

import (
	"fetch-app/handlers"
	"fetch-app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func router(c handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/fetch_list", c.List).Methods(http.MethodGet)
	r.HandleFunc("/fetch_list_usd", c.ListUSD).Methods(http.MethodGet)
	r.HandleFunc("/aggregates", c.Aggregate).Methods(http.MethodGet)
	r.Use(middlewares.AuthRequired)
	return r
}
