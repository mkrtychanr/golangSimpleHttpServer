package server

import (
	"github.com/mkrtychanr/wildberries_L0/internal/database"
	"github.com/mkrtychanr/wildberries_L0/internal/model"
)

type Server struct {
	cache  map[string]model.Order
	db     *database.Database
	config *config
}

func setConfigs(path string) (*database.Database, *config, error) {
	db, err := database.SetConfig(path)
	if err != nil {
		return nil, nil, err
	}
	config, err := newConfig(path)
	if err != nil {
		return nil, nil, err
	}
	return db, config, nil
}

func NewServer(path string) (*Server, error) {
	db, config, err := setConfigs(path)
	if err != nil {
		return nil, err
	}
	return &Server{
		db:     db,
		cache:  make(map[string]model.Order),
		config: config,
	}, nil
}

func (s *Server) Up() {
	s.db.Open()
}

func (s *Server) Down() {
	s.db.Close()
}
