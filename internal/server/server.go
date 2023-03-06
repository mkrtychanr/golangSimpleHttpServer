package server

import (
	"github.com/mkrtychanr/wildberries_L0/internal/database"
	"github.com/mkrtychanr/wildberries_L0/internal/model"
)

type Server struct {
	cache map[string]model.Order
	db    *database.Database
}

func NewServer(path string) (*Server, error) {
	db, err := database.SetConfig(path)
	if err != nil {
		return nil, err
	}
	cache := make(map[string]model.Order)
	return &Server{
		db:    db,
		cache: cache,
	}, nil
}

func (s *Server) Up() {
	s.db.Open()
}

func (s *Server) Down() {
	s.db.Close()
}
