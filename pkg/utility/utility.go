package utility

import (
	"fmt"
)

func ArmstrongOrNot_1(num int) {
	// Copy the original number
	orgNum := num

	// Find the number of digits in the number
	var digits int
	for {
		if num == 0 {
			break
		}
		num /= 10
		digits += 1
	}

	// Reset the value of num
	num = orgNum

	// find the sum
	var sumTotal, rem, temp int

	for {
		// Loop ending condition
		if num == 0 {
			break
		}
		// Reset temp
		temp = 1
		// Find the last digit of num
		rem = num % 10
		// Update the num with the last digit
		num = num / 10
		// Loop for multipling each digit b number of times
		for i := 1; i <= digits; i++ {
			temp = temp * rem
		}
		sumTotal = sumTotal + (temp)
	}

	// Check for equality
	if orgNum == sumTotal {
		fmt.Printf("%d is an Armstrong number.", orgNum)
	} else {
		fmt.Printf("%d is not an Armstrong Number.", orgNum)
	}
}
