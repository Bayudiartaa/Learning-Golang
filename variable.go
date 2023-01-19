package main

import "fmt"

func main() {
	// menginisiasi type data 
	var name string
	
	name = "ama"
	fmt.Println(name)

	// tidak menginisiasi type data 
	var age = 100
	fmt.Println(age) 

	// pengganti var tapi dgn variable yang sama
	country := "Indonesian"
	fmt.Println(country)

	country = "USA"
	fmt.Println(country)

	// deklaration multiple variable

	var (
		firstName = "Ama"
		lastName = "Lasna"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)


}