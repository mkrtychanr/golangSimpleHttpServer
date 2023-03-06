package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mkrtychanr/wildberries_L0/internal/server"
)

func HandleInterrupt(s *server.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("\r")
		s.Down()
		os.Exit(0)
	}()
}

func main() {
	server, err := server.NewServer("config.json")
	if err != nil {
		panic(err)
	}
	HandleInterrupt(server)
	server.Up()
	for {
	}
}
