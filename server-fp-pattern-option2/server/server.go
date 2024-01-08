package server

import "time"
import "rsc.io/quote"
import "fmt"

type Server struct {
  cfg Config
}

type Config struct {
  Host    string
  Port    int
  Timeout time.Duration
  MaxConn int
}

func New(cfg Config) *Server {
  return &Server{cfg}
}

func (server *Server) Hello() string {
  return quote.Hello()
}

func (s *Server) Start() error {
  fmt.Printf("My quote %v", quote.Hello())
  return nil
}
