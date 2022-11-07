package main

import "fmt"

var msg string

func updateMessage(s string) {
	msg = s
}

func printSomething() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world!"

	updateMessage("Hello, Universe!")
	printSomething()
}
