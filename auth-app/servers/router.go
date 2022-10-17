package servers

import (
	"auth-app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func router(c handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/info", c.Info).Methods(http.MethodGet)
	r.HandleFunc("/login", c.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", c.Register).Methods(http.MethodPost)
	return r
}
