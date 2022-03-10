package main

import "fmt"

// intger array type
type IntArray []int

// linear-search method
func (a IntArray) linearSearch(key int) (bool, int) {
	for i, element := range a {
		if element == key {
			return true, i
		}
	}
	return false, -1
}

// main
func main() {
	// program to do linear search using methods
	var n, key int
	fmt.Printf("Enter the size of the array: ")
	fmt.Scanf("%d", &n)

	arr := make(IntArray, n)

	fmt.Printf("Enter the array elements: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Printf("Enter the key: ")
	fmt.Scanf("%d", &key)
	found, index := arr.linearSearch(key)

	if found {
		fmt.Printf("Element found at index %v\n", index)
	} else {
		fmt.Printf("Element not found!!!\n")
	}
}
