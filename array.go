package main

import "fmt"

func main() {
	// set nilai array nya
	var books [3]string
	books[0] = "Vision"
	books[1] = "Filosofi Teras"
	books[2] = "Atomic habbit"

	fmt.Println(books[0])
	fmt.Println(books[1])
	fmt.Println(books[2])

	var books1 = [3]string {
		"Atomic habbit",
		"Filosofi teras",
		"Filos",
	}

	fmt.Println(books1) // [Atomic habbit Filosofi teras Filos]
	fmt.Println(books1[0]) // Atomic habbit
	fmt.Println(books1[1]) // Filosofi teras
	fmt.Println(books1[2]) // Filos
	// menghitung jumlah panjang data array
	fmt.Println(len(books1)) // 3
}