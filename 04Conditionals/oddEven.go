package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num)
	if num%2 == 0 {
		fmt.Printf("The number %v is even!!!", num)
	} else {
		fmt.Printf("The number %v is odd!!!", num)
	}
}
