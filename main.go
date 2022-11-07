package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	// lock resource
	m.Lock()
	msg = s
	// Unlock resource
	m.Unlock()
}

func printSomething() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, Universe!", &wg, &mutex)
	go updateMessage("Hello, Universe!", &wg, &mutex)
	wg.Wait()
	printSomething()
}
