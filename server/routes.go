package server

func (s *Server) routes() {

	s.router.HandleFunc("/", s.handleIndex())
}
