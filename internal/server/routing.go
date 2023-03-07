package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func (s *Server) startRouting() {
	s.router.Use(middleware.Logger)
	s.router.Get("/order/{order_uid}", s.handleGetId)
	address := fmt.Sprintf("%s%s", s.config.Host, s.config.Port)
	str := fmt.Sprintf("Server is up on %s%s", s.config.Host, s.config.Port)
	logrus.Info(str)
	http.ListenAndServe(address, s.router)
}

func (s *Server) handleGetId(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "order_uid")
	str := fmt.Sprintf("Someone requested %s, %s", id, request.Method)
	logrus.Info(str)
	data, ok := s.cache[id]
	if !ok {
		writer.Write([]byte("Something went wrong"))
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	writer.Write(b)
}
