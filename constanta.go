package main

import "fmt"

func main() {
	// variable yang tidak dapat diubah nilainya
	const firstName string = "Ama"
	const lastName = "Jani"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// deklaration multiple const
	const (
		age = 100
	    city = "tulungagung"
	)

	fmt.Println(age)
	fmt.Println(city)
}