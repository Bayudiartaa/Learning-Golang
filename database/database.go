package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init executed")
	connection = "PostgreSQL"
}

func GetDatabase() string {
	return connection
}