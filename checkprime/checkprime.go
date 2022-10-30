package checkprime

// Simplest form of Checking Prime number
func CheckPrime(num int) bool {
	var isPrime bool = true

	for i := 2; i <= (num/2)+1; i++ {
		if num%i != 0 {
			continue
		} else {
			isPrime = false
			break
		}
	}

	return isPrime
}

// Complex One than CheckPrime()
// Checks for corner cases
func CheckPrime_Method2(num int) bool {
	// Check valid input
	if num < 1 {
		panic("Entered number should be greater than 1")
	}

	// numbers from 1 - 3 are prime
	if (num >= 1) && (num <= 3) {
		return true
	}

	// numbers divisible by 2 or 3 are not prime
	if (num%2 == 0) || (num%3 == 0) {
		return false
	}

	// If a number is not prime it should be divisible by a number till its square root.
	// Also check the number is divided by 5 or 7
	for i := 5; i*i < num; i += 6 {
		if (num%i == 0) || (num%i+2 == 0) {
			return false
		}
	}

	return true
}
