package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	x = math.Sqrt(x)
	return x
}

func isPrime(n int) (isP bool) {
	k := int(Sqrt(float64(n)))
	for i := 2; i <= k; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true

}

func getListOfPrimes(n int) (gLOP []int) {
	x := make([]int, n)
	x[0] = 2
	counter := 1
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			x[counter] = i
			counter += 1
		}
	}
	return x[:counter]

}

func main() {
	number := int64(600851475143)
	sN := int(Sqrt(float64(number)))
	s := int(Sqrt(float64(sN)))
	y := make([]int, s)
	x := getListOfPrimes(sN)
	counter := 0
	for i := 0; i < len(x); i++ {
		if int64(number)%int64(x[i]) == 0 {
			y[counter] = x[i]
			counter += 1
		}

	}
	fmt.Println(y[:counter])
}