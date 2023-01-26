package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
   firstName = "Ama"
   middleName = "Jani"
   lastName = "Ini"
   return
}

func main() {
    a, b, c := getFullName()
	fmt.Println(a, b, c)
}