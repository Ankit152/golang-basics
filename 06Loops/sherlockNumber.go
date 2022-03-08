/*
Program to find number of sherlock numbers in a given range.
Sherlock number is a number whose sum of the digits is even.
*/
package main

import "fmt"

// checks if the number is even or not
func isEve(num int) bool {
	if num&1 == 0 {
		return true
	}
	return false
}

// checks if the number is shelock or not
func isSherlock(num int) bool {
	sum := 0
	for num > 0 {
		sum += (num % 10)
		num /= 10
	}
	return isEve(sum)
}

// calculate sherlock number in a given range
func calculateSherlock(num1 int, num2 int) int {
	count := 0
	for i := num1; i <= num2; i++ {
		if isSherlock(i) {
			count++
		}
	}
	return count
}

// main
func main() {
	var num1, num2 int
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &num2)

	fmt.Println("Total number of Sherlock nunber is: ", calculateSherlock(num1, num2))
}
