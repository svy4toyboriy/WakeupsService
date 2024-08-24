package server

import (
	"context"
	"encoding/json"
	"net/http"

	"WebService/config"
	"WebService/db"
	"github.com/gorilla/mux"
)

type Server struct {
	context context.Context
	config  *config.Config
	router  *mux.Router
	db      *db.PgDB
}

func Init(ctx context.Context, config *config.Config, db *db.PgDB) *Server {
	router := mux.NewRouter()
	s := &Server{
		context: ctx,
		config:  config,
		router:  router,
		db:      db,
	}
	s.routes()
	return s
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			// TODOO
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
