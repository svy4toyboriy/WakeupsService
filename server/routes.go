package server

func (s *Server) routes() {

	s.router.HandleFunc("/", s.handleIndex())

	s.router.HandleFunc("/schedule", s.handleCallSchedule()).Methods("POST")
}
