package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on structs")
	//no inheritance in golang; No super or parent

	fahim := User{"Fahim", "fahimanjon11710072@gmail.com", true, 24}
	fmt.Println(fahim)                                                     //oneline output (only values)
	fmt.Printf("fahim's details are: %+v\n", fahim)                        //detailed output "%+v" (both values and structures)
	fmt.Printf("Name is: %v and Email is: %v.\n", fahim.Name, fahim.Email) //Finding specific values
	fahim.GetStatus()
	fahim.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "xyz@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
