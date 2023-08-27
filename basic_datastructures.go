package main

import "fmt"

func main() {
	// Creating a slice
	slice := []int{1, 2, 3, 4, 5}

	// Getting length of slice
	length := len(slice)

	// Getting a sub-slice
	subSlice := slice[1:3]

	// Appending to slice
	slice = append(slice, 6)

	fmt.Println(slice)
	fmt.Println(length)
	fmt.Println(subSlice)

	// Creating a map
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Accessing a value by key
	value := myMap["a"]
	fmt.Println(value)

	// Adding a key-value pair
	myMap["d"] = 4

	// Deleting a key-value pair
	delete(myMap, "b")

	fmt.Println(myMap)

}
