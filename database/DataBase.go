package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	"os"
)

type DataBase struct {
	Db *pg.DB
}

func NewDataBase() *DataBase {
	db := pg.Connect(&pg.Options{
		User:     "root",
		Password: "root_mpt",
		Database: "jg",
		Addr:     "194.169.163.29:5432",
	})
	if db.Ping(context.Background()) != nil {
		os.Exit(1)
	}

	return &DataBase{Db: db}
}

func (b DataBase) Insert(s interface{}) {
	_, err := b.Db.Model(s).Insert()
	if err != nil {
		panic(err)
	}
}
