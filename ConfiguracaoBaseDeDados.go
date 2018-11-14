package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "usr_app"
	dbpass = "123456"
	dbname = "usr_app"
)

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
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
