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
