package main

import (
	"fmt"
	"strings"
)

// vardiac function that returns the joined string
func joinString(elements ...string) string {
	return strings.Join(elements, "-")
}

// main
func main() {
	// program to join elements of a slice using vardiac functions
	elements := []string{
		"hello", "hola", "namaste",
		"blah", "blah", "blah"}

	fmt.Println(joinString(elements...)) // slice '...'
}
