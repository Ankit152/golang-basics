package main

import "fmt"

func main() {
	defer fmt.Println("from Golang!")
	defer fmt.Println("World")
	fmt.Println("Hello")
}
