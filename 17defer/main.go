package main

import "fmt"

func main() {
	defer fmt.Println("Welcome fourth time")
	defer fmt.Println("Welcome third time")
	defer fmt.Println("Welcome second time")
	fmt.Println("Welcome to class on defer")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v\n", i)
	}
}

//output will be like this--
// Welcome to class on defer
// 4
// 3
// 2
// 1
// 0
// Welcome second time
// Welcome third time
// Welcome fourth time
// so, you understand the order of output? it's because of keyword "defer"
