package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printSomething() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, Universe!", &wg)
	go updateMessage("Hello, Universe!", &wg)
	wg.Wait()
	printSomething()
}
