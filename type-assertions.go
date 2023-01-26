package main 

import "fmt"

func Random() interface{} {
	return true
}

func main() {
	var result interface{} = Random()
    // resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Uknown")
	}
}