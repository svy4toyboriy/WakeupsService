package server

import (
	"net/http"
)

func (s *Server) handleCallSchedule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.ScheduleCall(w, r)
	}
}
