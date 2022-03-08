package main

import "fmt"

// function to check if a number is power of two or not and return accordingly
func isPowTwo(num int) bool {
	// conditions check
	if num&(num-1) == 0 {
		// return value
		return true
	}
	// return value
	return false
}

// main
func main() {
	// program to check if a number is a power of two using fucntions
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	// conditions check
	if isPowTwo(num) == true {
		fmt.Println("The number", num, "is a power of 2!!!")
	} else {
		fmt.Println("The number", num, "is not a power of 2!!!")
	}
}
