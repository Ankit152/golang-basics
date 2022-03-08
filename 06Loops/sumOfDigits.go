package main

import "fmt"

// function to calculate sum of the digits of a number
func sumOfDigits(num int) {
	sum := 0
	for num > 0 {
		sum += (num % 10)
		num /= 10
	}
	fmt.Println("The sum of the digits of the given number is:", sum)
}

// main
func main() {
	// program to find the sum of the digits of a number
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num)

	// function call
	sumOfDigits(num)
}
