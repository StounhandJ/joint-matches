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

func (b DataBase) Insert(s struct{}) {
	field, marks := getMarksValues(s)
	b.db.Exec(fmt.Sprintf("insert into %s(%s) values(%s)", reflect.TypeOf(s).Name(), field, marks))
}

func getMarksValues(s struct{}) (string, string) {
	t := reflect.TypeOf(s)
	count := 0
	var str1, str2 string
	for i := 0; i < t.NumField(); i++ {
		v1, v2 := t.Field(i).Tag.Lookup("pg")
		if v2 {
			str1 += fmt.Sprintf("%s, ", v1)
			str2 += fmt.Sprintf("$%s, ", v1)
			count += 1
		}
	}
	if count == 0 {
		return "", ""
	}
	return str1[0 : len(str1)-2], str2[0 : len(str2)-2]
}
