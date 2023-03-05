package migrate

import (
	"github.com/go-pg/pg/orm"
	"github.com/mkrtychanr/wildberries_L0/internal/database"
	"github.com/mkrtychanr/wildberries_L0/internal/model"
)

func CreateSchema(db *database.Database) error {
	models := []interface{}{
		(*model.Order)(nil),
	}

	for _, model := range models {
		op := orm.CreateTableOptions{}
		err := db.DB.Model(model).CreateTable(&op)
		if err != nil {
			return err
		}
	}
	return nil
}
