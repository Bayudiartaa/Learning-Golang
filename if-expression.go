package main

import "fmt"

func main() {
	name :=  "Ama"

	if name == "Ama" {
		fmt.Println("hai Ama")
	} else if name == "ani" {
		fmt.Println("Hai ani")
	} else {
		fmt.Println("Kamu salah manggil namaku :(")
	}


	if length := len(name);length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Sip Udah benar")
	}
}clea