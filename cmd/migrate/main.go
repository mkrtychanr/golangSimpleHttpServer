package main

import (
	"github.com/mkrtychanr/wildberries_L0/internal/database"
	"github.com/mkrtychanr/wildberries_L0/internal/migrate"
)

func main() {
	db, err := database.SetConfig("config.json")
	if err != nil {
		panic(err)
	}
	db.Open()
	defer db.Close()
	if err = migrate.CreateSchema(db); err != nil {
		panic(err)
	}
}
