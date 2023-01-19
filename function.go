package main

import "fmt"


func sayHelloTo(firstName string, secondName string) {
	fmt.Println("Hello " + firstName + " " + secondName)
}

func getName(name string) string {
	return "Hai" + name
}

func getData() (string, int) {
	return "AMA", 120
}

func main() {
	sayHelloTo("AMA", "Jani")
	result := getName("SALMA")
	fmt.Println(result)

	// name, age := getData()
	// fmt.Println(name, age)

	name, _ := getData()
	fmt.Println(name)
}