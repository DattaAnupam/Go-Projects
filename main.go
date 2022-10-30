package main

import (
	"fmt"

	"github.com/anupam/prime-not-prime/checkprime"
)

func main() {
	// Take input from user
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

	// check Prime or not
	isPrime := checkprime.CheckPrime(num)
	// if prime print message
	if isPrime {
		fmt.Printf("%d is a Prime number", num)
	} else {
		fmt.Printf("%d is not a Prime number", num)
	}
}
