package main

import (
	"fmt"
	"sync"
)

func PrintSomething(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(name)
}

func main() {
	var wg sync.WaitGroup

	names := []string{
		"Anupam",
		"Bapi",
		"Tapas",
		"Roxy",
		"Shubhrojit",
		"Abhijit",
	}

	// Add go routine counts
	wg.Add(len(names))

	for i, name := range names {
		go PrintSomething(fmt.Sprintf("%d: %s", i, name), &wg)
	}
	// Wait until all go routines are performed
	wg.Wait()

	wg.Add(1)
	go PrintSomething("Final Name", &wg)
	wg.Wait()
}
