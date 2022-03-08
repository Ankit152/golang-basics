package main

import "fmt"

// main
func main() {
	// program to iterate over a given range of numbers
	var num1, num2 int
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &num2)

	// iterate using loops single line
	for i := num1; i <= num2; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// iterate using multiple lines (while loop like structure)
	i := num1
	for i <= num2 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

}
