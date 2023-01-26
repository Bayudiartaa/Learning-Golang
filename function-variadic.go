package main

import "fmt"

func sumAll(numbers ...int) int{
    total := 0

	for _, value := range numbers {
        total += value
	}
	return total
}

func main() {
   total := sumAll(10,10,10)
   fmt.Println(total) //30

   //variadic function with slice of numbers 
   slice := []int{10,20,30,40}
   total = sumAll(slice...)
   fmt.Println(total) //100
}