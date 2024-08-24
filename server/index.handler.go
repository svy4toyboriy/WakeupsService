package server

import (
	"net/http"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, map[string]interface{}{
			"version": "1.0",
			"name":    "My dope web service is now Running",
		}, 200)
	}
}
