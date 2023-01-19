package main

import "fmt"

func main() {
	
	for counter := 1; counter <= 10; counter-=-1 {
		fmt.Println("loop to", counter)
		
	}

	slice := []int{1,2,3,4,5,5,6,7,8,8}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Eka"
	person["address"] = "kediri"

	for key, value := range person {
		fmt.Println(key, "=", value) 
		//name = Eka
        //address = kediri
	}
}