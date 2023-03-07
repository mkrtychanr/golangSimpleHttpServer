package database

import (
	"github.com/go-pg/pg"
	"github.com/mkrtychanr/wildberries_L0/internal/model"
)

type Database struct {
	DB     *pg.DB
	config *config
}

func SetConfig(path string) (*Database, error) {
	config, err := newConfig(path)
	if err != nil {
		return nil, err
	}
	return &Database{nil, config}, nil
}

func (db *Database) Open() {
	db.DB = pg.Connect(&pg.Options{
		User:     db.config.User,
		Password: db.config.Password,
		Database: db.config.Database,
	})
}

func (db *Database) Close() {
	db.DB.Close()
}

func (db *Database) AddOrder(order model.Order) error {
	if _, err := db.DB.Model(&order).Insert(); err != nil {
		return err
	}
	return nil
}
