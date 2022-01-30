package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)
	if num&(num-1) == 0 {
		fmt.Printf("The number %v is a power of two!!!", num)
	} else {
		fmt.Printf("The number %v is not a power of two!!!", num)
	}
}
