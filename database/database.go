package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

var defaultConnection = "host=127.0.0.1 port=5432 user=postgres dbname=sched password=postgres sslmode=disable"

func Connect(connString string) *gorm.DB {
	if (connString == "") {
		fmt.Println("No connection string specified, using the default local connection.")
		connString = defaultConnection
	}

	conn, err := gorm.Open("postgres", connString)
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

