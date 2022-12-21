package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") //adding in array
	fmt.Println(fruitList)

	//slicing [1(start from index):3(stop before index)]. Meaning last ranges are not inclusive
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //will print "[Tomato Peach Mango Banana]"

	//now previous output became the default list
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) //will operate on the updated list and will print "[Peach Mango]"

	//again previous output became the default list
	fruitList = append(fruitList[:1])
	fmt.Println(fruitList) //will operate on the updated list and will print "[Peach]"

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777, ouput will return an error as size was allocated 4 already

	highScores = append(highScores, 555, 666, 321) //but this time output will not return an error though size is bounded because of make() method

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
