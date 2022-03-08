package main

import "fmt"

// main
func main() {
	// program to perform all assignment operations
	var num1, num2 int
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &num2)

	// simple assignment
	num1 = num2
	fmt.Println("Simple assignment(=): ", num1)

	// add assignmnet
	num1 += num2
	fmt.Println("Add assignment(+=): ", num1)

	// subtract assignment
	num1 -= num2
	fmt.Println("Subtract assignment(-=): ", num1)

	// multiply assignment
	num1 *= num2
	fmt.Println("Multiply assignment(*=): ", num1)

	// division assignment
	num1 /= num2
	fmt.Println("Division assignment(/=): ", num1)

	// modulus assignment
	num1 %= num2
	fmt.Println("Modulus assignment(%=): ", num1)
}
