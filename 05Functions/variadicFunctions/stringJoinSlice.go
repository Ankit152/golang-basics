package main

import (
	"fmt"
	"strings"
)

// vardiac function that returns the joined string
func joinString(str1 string, elements ...string) string {
	str1 = strings.Replace(str1, " ", "-", -1)
	return str1 + "-" + strings.Join(elements, "-")
}

// main
func main() {
	// program to join elements of a slice using vardiac functions
	elements := []string{
		"hello", "hola", "namaste",
		"blah", "blah", "blah"}
	str1 := "Hi there"
	fmt.Println(joinString(str1, elements...)) // slice '...'
}
