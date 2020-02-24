package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

func Connect() *gorm.DB {
	conn, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=sched password=postgres sslmode=disable")
	if err != nil {
		fmt.Println("db err: ", err)
		panic("Could not connect to db")
	}
	connection = conn
	return conn
}

func Get() *gorm.DB {
	return connection
}

