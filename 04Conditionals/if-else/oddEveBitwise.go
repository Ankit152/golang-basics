package main

import "fmt"

// main
func main() {
	// program to check odd even number using bitwise operator
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)

	// condition check
	if num&1 == 1 {
		fmt.Println("The entered number is Odd!!")
	} else {
		fmt.Println("The entered number is Even!!")
	}
}
