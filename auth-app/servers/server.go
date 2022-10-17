package servers

import (
	"auth-app/databases"
	"auth-app/handlers"
	"context"
	"log"
	"net/http"
)

// Server struct
type Server struct {
	Hndl *handlers.Handler
}

// Init func
func Init() (*Server, error) {
	s := Server{Hndl: &handlers.Handler{}}
	ctx := context.Background()
	db, err := databases.Connect(
		ctx,
		"mongodb+srv://admin:19101998@mongodbcloud.mpbrizl.mongodb.net/?retryWrites=true&w=majority",
		"efishery",
		"user",
	)
	if err != nil {
		return &s, err
	}

	s.Hndl.Context = ctx
	s.Hndl.Database = db
	return &s, nil
}

// Start func
func (s *Server) Start() error {
	r := router(*s.Hndl)
	log.Println("server starts on", ":3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}
	return nil
}
