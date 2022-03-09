package main

import "fmt"

// fucntion to check if it is odd or not
func isOdd(num int) bool {
	// condition check
	if num&1 == 1 {
		// value return
		return true
	}
	// value return
	return false
}

// main
func main() {
	// program to check if a number is odd/eve using fucntions
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	// conditions check
	if isOdd(num) {
		fmt.Println("The number", num, "is odd!")
	} else {
		fmt.Println("The number", num, "is even!")
	}
}
