package main

import "fmt"

func main() {
	fmt.Println("This is all about operators in Golang!!!")
	var num1, num2 int
	fmt.Printf("Enter two numbers: ")
	fmt.Scanf("%d %d", &num1, &num2)
	add := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	fmt.Println("Addition of two number is: ", add)
	fmt.Println("Subtraction of two number is: ", sub)
	fmt.Println("Multiplication of two number is: ", mul)
	fmt.Println("Division of two number is: ", div)
}
