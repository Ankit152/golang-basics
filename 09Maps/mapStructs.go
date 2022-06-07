package main

import "fmt"

func main() {
	enum := map[string]struct{}{
		"one":   {},
		"two":   {},
		"three": {},
	}

	for key, val := range enum {
		fmt.Printf("Key: %v -- Value: %v\n", key, val)
	}
}
