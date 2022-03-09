package main

import "fmt"

// struct books
type books struct {
	price, quantity, code int
}

// main
func main() {
	// program to iterate over struct

	// array/slice of structs
	bookShelf := []books{
		{
			price:    200,
			quantity: 20,
			code:     1,
		},
		{
			price:    250,
			quantity: 32,
			code:     2,
		},
		{
			price:    120,
			quantity: 10,
			code:     3,
		},
		{
			price:    150,
			quantity: 22,
			code:     4,
		},
		{
			price:    220,
			quantity: 40,
			code:     5,
		},
		{
			price:    500,
			quantity: 2,
			code:     6,
		},
	}

	fmt.Printf("List of books: %d\n", len(bookShelf))
	for i, book := range bookShelf {
		fmt.Printf("Book id is %d and book price is: %d\n", i, book.price)
	}

	// count the list of books whose price is b/w 200-300 and quantity less than 10
	count := 0
	for _, book := range bookShelf {
		if book.price >= 200 && book.price <= 300 && book.quantity >= 10 {
			count++
		}
	}
	fmt.Printf("The list of books whose price is b/w 200-300 and quantity greater than 10 is: %d\n", count)
}
