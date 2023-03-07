package main

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	path := ""
	fmt.Printf("Enter the path to file: ")
	fmt.Scan(&path)
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	sc.Publish("addNewOrder", b)
}
