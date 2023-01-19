package main

import "fmt"

func main() {

	var name1 = "Ama"
	var name2 = "Ama"

	var result bool = name1 == name2
	fmt.Println(result) // true
    
	var name3 = "Ama"
	var name4 = "jani"

	var result1 bool = name3 == name4
	fmt.Println(result1) // false


}