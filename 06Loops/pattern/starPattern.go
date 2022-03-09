package main

import "fmt"

// pyramid function
func pyramid(n int) {
	/*
		Output
		*
		* *
		* * *
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// pyramid2 function
func pyramid2(n int) {
	/*
		Output
		* * *
		* *
		*
	*/
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// pyramid3 function
func pyramid3(n int) {
	/*
		Output
			*
		   * *
		  * * *
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// parallelogram function
func parallelogram(n int) {
	/*
		Output
			* * *
		   * * *
		  * * *
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// square function
func square(n int) {
	/*
		Output
		* * *
		* * *
		* * *
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// kite function
func kite(n int) {
	/*
		Output
			*
		   * *
		  * * *
		   * *
		    *
	*/
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := 1; i <= n-1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= n-i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// main
func main() {
	// program to do some old school star patterns in golang!
	var n, choice int
	fmt.Print("Enter the max row size: ")
	fmt.Scanf("%d", &n)
	fmt.Print("1. Pyramid-I\n2. Pyramid-II\n3. Pyramid-III\n4. Parallelogram\n5. Square\n6. Kite\n")
	fmt.Print("Enter the choice: ")
	fmt.Scanf("%d", &choice)

	// switch cases
	switch choice {
	case 1:
		pyramid(n)
	case 2:
		pyramid2(n)
	case 3:
		pyramid3(n)
	case 4:
		parallelogram(n)
	case 5:
		square(n)
	case 6:
		kite(n)
	default:
		fmt.Println("Invalid choice. Wrong Input!!!")
	}
}
