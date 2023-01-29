package main

import (
	"fmt"
	"github.com/Bayudiartaa/Learning-Golang/database"
	// _ "github.com/Bayudiartaa/Learning-Golang/database"
)

func main() {
	result := database.GetDatabase
	fmt.Println(result)
}