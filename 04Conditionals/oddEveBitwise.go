package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &num)
	if num&1 == 1 {
		fmt.Println("The entered number is Odd!!")
	} else {
		fmt.Println("The entered number is Even!!")
	}
}
