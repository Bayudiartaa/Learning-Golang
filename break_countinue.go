package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
	    }
		fmt.Println("loop to", i)
    }

	for i := 0; i < 10; i-=-1 {
		if i%2 == 0 {
			continue
		}
		fmt.Println("loop to", i)
	}
}