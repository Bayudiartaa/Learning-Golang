package main

import "fmt"

func main() {
	person := map[string]string{
		"name" : "ama",
		"address" : "jepu",
	}
    //mengudah data map dgn key 
	person["title"] = "Mixue"

	fmt.Println(person) //map[address:jepu name:ama title:Mixue]
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

    // make map
	var book map[string]string = make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "amay"
	book["address"] = "sukabumi"
	//before delete
	fmt.Println(book) // map[address:sukabumi author:amay title:Golang]

	delete(book, "title")
    //after delete
	fmt.Println(book) //map[address:sukabumi author:amay]
}