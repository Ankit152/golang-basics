package main

import (
	"fmt"
	"strings"
)

// vardiac function
func joinString(str ...string) string {
	return strings.Join(str, "-")
}

// main
func main() {
	// program to join string with variadic functions
	fmt.Println(joinString("Hi there!", "Golang", "Tutorial"))
	fmt.Println(joinString("Alpha", "to", "delta", "charlie", "tango"))
}
