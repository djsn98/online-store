package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Address string
	Router  *mux.Router
}

func (s *Server) Init(address string, router *mux.Router) *Server {
	s.Address = address
	s.Router = router
	return s
}

func (s Server) Run() {
	fmt.Printf("server started at %s\n", s.Address)
	err := http.ListenAndServe(s.Address, s.Router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
