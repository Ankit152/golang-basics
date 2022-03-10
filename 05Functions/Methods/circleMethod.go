package main

import "fmt"

// circle struct
type Circle struct {
	radius float64
}

// area method
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// preimeter method
func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// main
func main() {
	// program to use methods
	var c Circle
	fmt.Printf("Enter the radius: ")
	fmt.Scanf("%f", &c.radius)

	fmt.Printf("Area of the circle is: %v\n", c.area())
	fmt.Printf("Perimeter of the circle is: %v\n", c.perimeter())
}
