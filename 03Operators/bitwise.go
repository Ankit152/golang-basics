package main

import "fmt"

func main() {
	fmt.Println("This ia all about bitwise operators!")
	var num1, num2 int
	fmt.Printf("Enter the two numbers: ")
	fmt.Scanf("%d %d", &num1, &num2)
	and := num1 & num2
	xor := num1 ^ num2
	or := num1 | num2
	leftShift := num1 << 3
	rightShift := num1 >> 1
	fmt.Println("The and of the two numbers is: ", and)
	fmt.Println("The or of the two numbers is: ", or)
	fmt.Println("The xor of the two numbers is: ", xor)
	fmt.Println("The 3 left shift of the number is: ", leftShift)
	fmt.Println("The 1 right shift of the number is: ", rightShift)
}
