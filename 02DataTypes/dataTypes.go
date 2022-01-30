package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// default data type declaration
	fmt.Println("This code is all about data types in Golang.")
	var num int = 64
	fmt.Printf("This is an %T type with the value %v and the size %v\n", num, num, unsafe.Sizeof(num))
	var choice bool = true
	fmt.Printf("This is an %T type with the value %v and the size %v\n", choice, choice, unsafe.Sizeof(choice))
	var name string = "Ankit Kurmi"
	fmt.Printf("This is an %T type with the value %v and the size %v\n", name, name, unsafe.Sizeof(name))
	var percent float32 = 234.54678916
	fmt.Printf("This is an %T type with the value %v and the size %v\n", percent, percent, unsafe.Sizeof(percent))
	var percent2 float64 = 234.54678916
	fmt.Printf("This is an %T type with the value %v and the size %v\n", percent2, percent2, unsafe.Sizeof(percent2))

	// walrus operator -- explicitly data type assignment
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("This is by using walrus operator --> :=")
	name2 := "theankitkurmi"
	fmt.Printf("This is an %T type with the value %v and the size %v\n", name2, name2, unsafe.Sizeof(name2))
	num2 := 152
	fmt.Printf("This is an %T type with the value %v and the size %v\n", num2, num2, unsafe.Sizeof(num2))

}
