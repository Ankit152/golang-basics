package main

import "fmt"

// linear search
func linearSearch(arr []int, key int) (bool, int) {
	// iteration over array
	for i, ele := range arr {
		if key == ele {
			return true, i
		}
	}
	return false, -1
}

// main
func main() {
	// linear search code
	var n, key int
	fmt.Printf("Enter the array size: ")
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	fmt.Printf("Enter the array elements: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Printf("Enter the key to be searched: ")
	fmt.Scanf("%d", &key)

	// function call
	found, index := linearSearch(arr, key)
	if found {
		fmt.Printf("Element found at index %v.\n", index)
	} else {
		fmt.Println("Not found.")
	}
}
