package main

import "fmt"

// pyramid function
func pyramid(n int) {
	/*
		Output
		1
		2 2
		3 3 3
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}

// pyramid2 fucntion
func pyramid2(n int) {
	/*
		Output
		1
		1 2
		1 2 3
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

// pyramid3 function
func pyramid3(n int) {
	/*
		Output
			1
		   2 2
		  3 3 3
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

// pyramid4 function
func pyramid4(n int) {
	/*
		Output
			1
		   1 2
		  1 2 3
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// main
func main() {
	var n, ch int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("1. Pyramid-I\n2. Pyramid-II\n3. Pyramid-III\n4. Pyramid-IV\n")
	fmt.Print("Enter the choice: ")
	fmt.Scanf("%d", &ch)

	// switch cases
	switch ch {
	case 1:
		pyramid(n)
	case 2:
		pyramid2(n)
	case 3:
		pyramid3(n)
	case 4:
		pyramid4(n)
	default:
		fmt.Println("Wrong Choice!!!")
	}
}
