package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("error: ", err.Error())
	}
}