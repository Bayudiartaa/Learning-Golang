package main

import "fmt"

func Ups (i int) interface{} {
	if i == 1 {
		return 1
	} else {
		return "UPS"
	}
}
func main() {
   kosong := Ups(2)
   fmt.Println(kosong)
}