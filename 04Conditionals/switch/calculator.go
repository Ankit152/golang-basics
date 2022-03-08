package main

import "fmt"

// main
func main() {
	// program to make a calculator using switch cases
	var num1, num2, choice int
	fmt.Println("####### Welcome to Golang Calculator #######")
	fmt.Println("List of operations that can be performed:")
	fmt.Print("1. ADD \n2. SUB \n3. MUL \n4. DIV\n")
	fmt.Printf("Enter the choice: ")
	fmt.Scanf("%d", &choice)
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the second number:  ")
	fmt.Scanf("%d", &num2)

	// switch case
	switch choice {
	case 1:
		fmt.Println("Addition is: ", num1+num2)
	case 2:
		fmt.Println("Subtraction is: ", num1-num2)
	case 3:
		fmt.Println("Multiplication is: ", num1*num2)
	case 4:
		// nested switch
		switch num2 {
		case 0:
			fmt.Println("Cannot divide by 0. Math error!!!")
		default:
			fmt.Println("Division is: ", num1/num2)
		}
	default:
		fmt.Println("Wrong Choice!!!")
	}
}
