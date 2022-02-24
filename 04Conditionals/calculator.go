package main

import "fmt"

func main() {
	var a, b, ch int
	fmt.Println("####### Welcome to Golang Calculator #######")
	fmt.Println("List of operations that can be performed:")
	fmt.Print("1. ADD \n2. SUB \n3. MUL \n4. DIV\n")
	fmt.Print("Enter 1st number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter 2nd number: ")
	fmt.Scanf("%d", &b)
	fmt.Print("Enter the choice: ")
	fmt.Scanf("%d", &ch)
	if ch == 1 {
		fmt.Println("Addition is:", a+b)
	} else if ch == 2 {
		fmt.Println("Subtraction is:", a-b)
	} else if ch == 3 {
		fmt.Println("Multiplication is:", a*b)
	} else if ch == 4 && b == 0 {
		fmt.Println("Math Error!!!")
	} else if ch == 4 {
		fmt.Println("Division is:", a/b)
	} else {
		fmt.Println("Not a valid operation!!!")
	}
}
