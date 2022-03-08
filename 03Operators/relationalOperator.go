package main

import "fmt"

// main
func main() {
	// program to perform all relational operation
	var num1, num2 int
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &num2)

	// relational operators
	fmt.Println(num1, ">", num2, "? :", num1 > num2)
	fmt.Println(num1, "<", num2, "? :", num1 < num2)
	fmt.Println(num1, ">=", num2, "? :", num1 >= num2)
	fmt.Println(num1, "<=", num2, "? :", num1 <= num2)
	fmt.Println(num1, "==", num2, "? :", num1 == num2)
	fmt.Println(num1, "!=", num2, "? :", num1 != num2)
}
