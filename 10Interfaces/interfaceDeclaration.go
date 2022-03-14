package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}
	printGreeting(eb)
	printGreeting(hb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	return "Hi there! Greetings from english bot."
}
func (hindiBot) getGreeting() string {
	return "Namaste! Hindi bot ki taraf se abhinandan."
}
