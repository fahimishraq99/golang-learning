package main

import "fmt"

const LoginToken string = "sadlife" //public

func main() {
	var username string = "fahim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //uint8 = 0 to 255//
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.746234793249279449
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style

	numberOfuser := 300000
	fmt.Println(numberOfuser)

	fmt.Println(LoginToken) //variable declaration from outside of a method
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
