package database

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
)

type DataBase struct {
	db *sql.DB
}

func NewDataBase() *DataBase {
	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &DataBase{db: db}
}

func (b DataBase) execute(s struct{}) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%+v\n", t.Field(i))
	}
}
