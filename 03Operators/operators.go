package main

import "fmt"

func main() {
	fmt.Println("This is all about operators in Golang!!!")
	var num1, num2 int
	fmt.Printf("Enter two numbers: ")
	fmt.Scanf("%d %d", &num1, &num2)
	add := num1 + num2
	sub := num1 - num2
	xor := num1 ^ num2
	and := num1 & num2
	or := num1 | num2
	mul := num1 * num2
	fmt.Println("Addition of two number is: ", add)
	fmt.Println("Subtraction of two number is: ", sub)
	fmt.Println("Xor of two number is: ", xor)
	fmt.Println("Or of two number is: ", or)
	fmt.Println("And of two number is: ", and)
	fmt.Println("Multiplication of two number is: ", mul)
}
