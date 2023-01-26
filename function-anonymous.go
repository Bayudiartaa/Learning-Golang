package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("You are welcome ", name)
	}
}

func main() {
   blacklist := func(name string) bool {
	   return name == "anjing"
   }

   registerUser("ama", blacklist)
   registerUser("anjing", blacklist)
   registerUser("anjing", func (name string) bool {
	return name == "anjas"
   })
}