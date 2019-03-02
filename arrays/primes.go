package arrays

// GeneratePrimes generates all prime numbers from 2
// all the way up but no including upperBound.
// Time complexity is O(n log(log(n))), space complexity O(n)
func GeneratePrimes(upperBound int) []uint {
	switch upperBound {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		return nil
	}

	isPrime := make([]bool, upperBound)
	for i := range isPrime {
		isPrime[i] = true
	}
	var primes []uint
	for i := 2; i < len(isPrime); i++ {
		if isPrime[i] {
			primes = append(primes, uint(i))
			for j := i; j < len(isPrime); j += i {
				isPrime[j] = false
			}
		}
	}
	return primes
}
