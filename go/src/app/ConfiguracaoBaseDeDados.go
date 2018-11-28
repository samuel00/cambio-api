package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

const (
	dbhost = "172.23.0.1"
	dbport = "5432"
	dbuser = "usr_cambio"
	dbpass = "123456"
	dbname = "usr_cambio"
)

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		db.SingularTable(true)
	}
	fmt.Println("Successfully connected to DataBase Postgres!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[dbhost] = dbhost
	conf[dbport] = dbport
	conf[dbuser] = dbuser
	conf[dbpass] = dbpass
	conf[dbname] = dbname
	return conf
}
