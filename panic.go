package main

import "fmt"

func endApp() {
	fmt.Println("Application end")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	} 
	fmt.Println("Aplikasi Berjalan")
}

func main() {
    runApp(false)
    runApp(true)
}