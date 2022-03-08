package main

import "fmt"

// main
func main() {
	// program to check if a number is a power of 2 or not
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	// condition check
	if num&(num-1) == 0 {
		fmt.Printf("The number %v is a power of two!!!", num)
	} else {
		fmt.Printf("The number %v is not a power of two!!!", num)
	}
}
