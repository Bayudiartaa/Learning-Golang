package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married(){
    man.Name = "Mr." + man.Name
}

func main(){
    ama := Man{"Ama"}
	ama.Married()
	fmt.Println(ama.Name)
}