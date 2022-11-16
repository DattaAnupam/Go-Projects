package main

import (
	"fmt"

	"github.com/anupam/armstrong/pkg/utility"
)

func main() {
	// Take user input
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scanln(&num)

	// Find Armstrong or not
	utility.ArmstrongOrNot_1(num)
}
