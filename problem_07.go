package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	sR := math.Sqrt(x)
	return sR
}

func isPrime(n int) (isP bool) {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	k := int(Sqrt(float64(n)))
	for i := 3; i <= k; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	Switch := 1	// switch for the for loop
	number := 1	// start with 1 and increment by 2 (i.e - 1, 3, 5, 7, 9, etc.)

	// start with 1 since we will not count 2, which is a prime
	primes := 1

	for Switch >= 1 {
		if isPrime(number) {	// check for primality if true, add 1 to prime counter
			primes += 1
		}
		if primes == 10001 {	// if prime counter == 10001, print that prime number
			fmt.Println(number)
			Switch = 0
		}
		number += 2
	}
}