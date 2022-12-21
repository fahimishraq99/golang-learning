package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on structs")
	//no inheritance in golang; No super or parent

	fahim := User{"Fahim", "fahimanjon11710072@gmail.com", true, 24}
	fmt.Println(fahim)                                                   //oneline output (only values)
	fmt.Printf("fahim's details are: %+v\n", fahim)                      //detailed output "%+v" (both values and structures)
	fmt.Printf("Name is: %v and Email is: %v.", fahim.Name, fahim.Email) //Finding specific values
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
