package servers

import (
	"context"
	"fetch-app/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

// Server struct
type Server struct {
	Hndl *handlers.Handler
}

// Init func
func Init() (*Server, error) {
	s := Server{Hndl: &handlers.Handler{}}
	ctx := context.Background()
	c := cache.New(30*time.Minute, 1*time.Hour)

	s.Hndl.Context = ctx
	s.Hndl.Cache = c
	return &s, nil
}

// Start func
func (s *Server) Start() error {
	r := router(*s.Hndl)
	log.Println("server starts on", os.Getenv("PORT"))
	if err := http.ListenAndServe(os.Getenv("PORT"), r); err != nil {
		return err
	}
	return nil
}
