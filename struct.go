package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHai(name string) {
    fmt.Println("Hello", name, "My Name is:", customer.Name)
}


func (a Customer) sayHuu(name string) {
    fmt.Println("Hello", name, "Hello :", a.Name)
}

func main() {
    var customer Customer
	customer.Name = "John"
	customer.Address = "Jakarta"
	customer.Age = 20

	customer.sayHai("ama")
	customer.sayHuu("joko")
	fmt.Println(customer)

	ama := Customer{
		Name: "ama",
		Address: "Jakarta",
		Age:20,
	}
	fmt.Println(ama)
}