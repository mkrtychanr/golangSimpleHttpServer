package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/mkrtychanr/wildberries_L0/internal/database"
	"github.com/mkrtychanr/wildberries_L0/internal/model"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cache  map[string]model.Order
	db     *database.Database
	config *config
	router *chi.Mux
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
		router: chi.NewRouter(),
	}, nil
}

func (s *Server) Up() {
	s.db.Open()
	logrus.Info("Database is up")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		s.Down()
		os.Exit(0)
	}()
	go s.startRouting()
}

func (s *Server) Down() {
	logrus.Info("Server is down")
	s.db.Close()
}
