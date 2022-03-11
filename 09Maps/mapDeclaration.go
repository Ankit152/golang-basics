package main

import "fmt"

// main
func main() {
	// program to declare/initialise and print maps
	countryCapital := map[string]string{
		"India":      "Delhi",
		"Sri Lanka":  "Colombo",
		"Bangladesh": "Dhaka",
		"Japan":      "Tokyo",
		"China":      "Beijing",
		"France":     "Paris",
	}

	fmt.Println("The country with their capitals are as follows: ")
	for country, capital := range countryCapital {
		// country is the key and city is the value
		fmt.Printf("The country is: %v and it's capital is: %v\n", country, capital)
	}

	// remove a country
	delete(countryCapital, "France")
	fmt.Printf("Deleted France from the key-value pair!\n")

	fmt.Println("The country with their capitals are as follows: ")
	for country, capital := range countryCapital {
		// country is the key and city is the value
		fmt.Printf("The country is: %v and it's capital is: %v\n", country, capital)
	}

	capital, ok := countryCapital["Italy"]
	if ok {
		fmt.Println("Italy's capital is present!", capital)
	} else {
		fmt.Println("Italy's capital is not present!")
	}
}
