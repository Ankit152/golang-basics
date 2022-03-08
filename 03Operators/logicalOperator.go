package main

import "fmt"

// main
func main() {
	// program to find max of three number
	var num1, num2, num3 int
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &num2)
	fmt.Printf("Enter the third number: ")
	fmt.Scanf("%d", &num3)

	var result int
	// conditions check
	if num1 > num2 && num1 > num3 {
		result = num1
	} else if num2 > num1 && num2 > num3 {
		result = num2
	} else {
		result = num3
	}

	fmt.Println("Max of all three is: ", result)
}
