package main

import "fmt"

// main
func main() {
	// program to check odd/even number
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num)

	// condition check
	if num%2 == 0 {
		fmt.Printf("The number %v is even!!!", num)
	} else {
		fmt.Printf("The number %v is odd!!!", num)
	}
}
