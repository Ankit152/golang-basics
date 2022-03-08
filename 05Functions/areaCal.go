package main

import "fmt"

// function to calculate area of a circle
func areaCricle() {
	var radius float64
	fmt.Print("Enter the radius: ")
	fmt.Scanf("%f", &radius)
	fmt.Println("Area of the circle is: ", 3.14*radius*radius)
}

// fucntion to calculate area of square
func areaSquare() {
	var length float64
	fmt.Print("Enter the side length: ")
	fmt.Scanf("%f", &length)
	fmt.Println("Area of the square is: ", length*length)
}

// function to calculate area of triangle
func areaTriangle() {
	var base, height float64
	fmt.Print("Enter the base: ")
	fmt.Scanf("%f", &base)
	fmt.Print("Enter the height: ")
	fmt.Scanf("%f", &height)
	fmt.Println("The area of the triangle is: ", 0.5*base*height)
}

// fucntion to calculate area of a rectangle
func areaRectangle() {
	var length, breadth float64
	fmt.Print("Enter the length: ")
	fmt.Scanf("%f", &length)
	fmt.Print("Enter the breadth: ")
	fmt.Scanf("%f", &breadth)
	fmt.Println("The area of the rectangle is: ", length*breadth)
}

// main
func main() {

	// declaration of choice
	var choice int
	fmt.Println("##### Program to calculate area #####")
	fmt.Println("1. Rectangle \n2. Circle\n3. Square\n4. Triangle")
	fmt.Print("Enter the choice: ")
	fmt.Scanf("%d", &choice)

	// call of methods based on choice
	if choice == 1 {
		areaRectangle()
	} else if choice == 2 {
		areaCricle()
	} else if choice == 3 {
		areaSquare()
	} else if choice == 4 {
		areaTriangle()
	} else {
		// wrong input message
		fmt.Println("Wrong Choice!!!")
	}
}
