package database

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
	"reflect"
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
	//field, marks := getMarksValues(s)
	//h := fmt.Sprintf("insert into %s(%s) values(%s)", reflect.TypeOf(s).Name(), field, marks)
	//fmt.Println(h)
	//_, err := b.db.Exec(fmt.Sprintf("insert into %s(%s) values(%s)", reflect.TypeOf(s).Name(), field, marks))
	//if err != nil {
	//	return
	//}
}

func getMarksValues(s interface{}) (string, string) {
	t := reflect.TypeOf(s)
	g := reflect.ValueOf(s)

	count := 0
	var str1, str2 string

	for i := 0; i < t.NumField(); i++ {
		v1, v2 := t.Field(i).Tag.Lookup("pg")
		if v2 {
			str1 += fmt.Sprintf("%s, ", v1)
			str2 += fmt.Sprintf("'%s', ", g.Field(i).String())
			count += 1
		}
	}
	if count == 0 {
		return "", ""
	}
	return str1[0 : len(str1)-2], str2[0 : len(str2)-2]
}
