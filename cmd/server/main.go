package main

import "github.com/mkrtychanr/wildberries_L0/internal/server"

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-chi/chi/v5"
// )

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello"))
// }

// func getID(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	_, err := strconv.Atoi(id)
// 	if err != nil {
// 		w.Write([]byte("Something went wrong"))
// 		return
// 	}
// 	w.Write([]byte(id))
// }

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", handleHello)
// 	r.Get("/{id}", getID)
// 	http.ListenAndServe(":1234", r)
// }

func main() {
	server, err := server.NewServer("config.json")
	if err != nil {
		panic(err)
	}
	server.Up()
	defer server.Down()
}
