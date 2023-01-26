package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
   goodBy := getGoodBye
   result := goodBy("caca")
   fmt.Println(result)
   fmt.Println(goodBy("ama"))
}