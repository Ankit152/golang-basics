package main

import "fmt"

// main
func main() {
	// program to check leap year
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanf("%d", &year)

	// condition check
	if year%400 == 0 {
		fmt.Println("It is a leap year!!!")
	} else if year%4 == 0 && year%100 != 0 {
		fmt.Println("It is a leap year!!!")
	} else {
		fmt.Println("It is not a leap year!!!")
	}
}
