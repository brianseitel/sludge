package server

import "net"

// Server ...
type Server struct {
	Listener net.Listener
	Down     bool
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		Down: true,
	}
}

// Start the server
func (s *Server) Start() {
	var err error
	s.Listener, err = net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		panic(err)
	}

	s.Down = false
}
