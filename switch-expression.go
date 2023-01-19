package main

import "fmt"

func main() {
	name := "ama"

	switch name {
	case "ama":
		fmt.Println("Hello ama")
	case "caca":
		fmt.Println("Hello caca")
	default:
		fmt.Println("Hello anon")
	}

	switch length := len(name);  length > 5 {
	case true:
		fmt.Println("Namaku Terlalu Panjang")
	case false:
		fmt.Println("namaku dah benar")
	} 

	length := len(name)
	switch {
    case length > 5:
		fmt.Println("Namaku Terlalu Panjang")
	case length < 2:
		fmt.Println("Namaku Terlalu Pendek")
	default:
		fmt.Println("Namaku udah benar")
	}
}