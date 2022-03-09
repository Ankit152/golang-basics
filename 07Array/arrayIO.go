package main

import "fmt"

// main
func main() {
	// program to do I/O in slices.
	var n int
	fmt.Printf("Enter the array size: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n) // used to define array of a partcular size

	// console input
	fmt.Printf("Enter the array elements(%v)\n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	// array output
	fmt.Printf("The array elements are: ")
	for _, ele := range arr {
		fmt.Printf("%v ", ele)
	}
	fmt.Println()
}
