package database

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
)

type DataBase struct {
	Db *pg.DB
}

func NewDataBase() *DataBase {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     os.Getenv("DB_ADDR"),
	})
	if db.Ping(context.Background()) != nil {
		fmt.Println("DataBase: no connection")
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
