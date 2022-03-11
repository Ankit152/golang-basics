// program to find the number of occurences of every element in an array
package main

import "fmt"

type IntArray []int

func main() {
	var n int
	fmt.Printf("Enter the size of the array: ")
	fmt.Scanf("%d", &n)
	arr := make(IntArray, n)
	fmt.Printf("Enter the array elements: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	freq := map[int]int{}
	for _, val := range arr {
		freq[val]++
	}
	fmt.Printf("The array elements are: ")
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	fmt.Println("Frequencies of each of the elements are as follows: ")
	for key, val := range freq {
		fmt.Printf("%v --> %v\n", key, val)
	}
}
