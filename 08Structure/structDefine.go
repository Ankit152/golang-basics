package main

import "fmt"

// defining a structure type
type student struct {
	name  string
	class int
	roll  int
	marks int // marks are out of 500
}

// main
func main() {
	var a student  // structure declaration
	fmt.Println(a) // blank, contains default values

	// directly assigning values to the structure variable
	a.name = "Ankit"
	a.class = 12
	a.roll = 4
	a.marks = 455
	fmt.Println(a)

	// naming the fields while initialising the structs
	s1 := student{
		name:  "Mayank",
		roll:  14,
		class: 12,
		marks: 470,
	}
	fmt.Println(s1)

	// naming the fields and leaving some of them
	s2 := student{
		name:  "Aditya",
		roll:  1,
		marks: 495}
	fmt.Println(s2)

	// passing the arguments in the order
	s3 := student{"Priyanshu", 25, 12, 460}
	fmt.Println(s3)

	s4 := student{"Yash", 50, 12, 470}
	fmt.Println(s4)

	s5 := student{"Ashish", 13, 12, 470}
	fmt.Println(s5)

	s6 := student{"Nico", 20, 12, 480}
	fmt.Println(s6)

	s7 := student{"Amador", 2, 12, 490}
	fmt.Println(s7)
}
