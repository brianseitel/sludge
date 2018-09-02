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
func (s *Server) Start(host string) error {
	var err error
	s.Listener, err = net.Listen("tcp", host)
	if err != nil {
		return err
	}

	s.Down = false
	return nil
}
