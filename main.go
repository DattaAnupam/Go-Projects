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
	isPrime, factors := checkprime.CheckPrime(&num)
	// isPrime := checkprime.CheckPrime_Method2(num)
	// if prime print message
	if isPrime {
		fmt.Printf("%d is a Prime number. \nOnly factors for %d are %v\n", num, num, factors)
	} else {
		fmt.Printf("%d is not a Prime number, and the factore(s) for it is/are: %v", num, factors)
	}
}
