package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on functions")
	greeter()

	result := adder(3, 5) //adding 2 values
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 5, 6, 2) //adding infinity values, also string
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
	//greeterTwo(); function inside function is not available
}

func adder(valOne int, valTwo int) int { //Type of the value should be mentioned outside of parenthesis according to what kind of value you want to be passed on. These are as known as 'function signatures'
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0 //storing the addition into 'total variable'

	for _, val := range values {
		total += val //every single time the value will be added with total, the value which will be given
	}

	return total, "Test Kore dekhchi"
}

// func greeterTwo() {
// 	fmt.Println("Another method")
// }

func greeter() { //Regarding sequence, call of functions does matter not the call of method
	fmt.Println("Hello from golang")
}
