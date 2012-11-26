package main

import "fmt"
import "math"

var Sum int64 = 0

func main() {
	
	counter := 1
	for counter < 2000000 {
		if isPrime(counter) {
			Sum += int64(counter)
		}
		counter += 2
	}
	fmt.Println(Sum + 2)

}

func Sqrt(x float64) float64 {
	x = math.Sqrt(x)
	return x
}

func isPrime(n int) (isP bool) {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	k := int(Sqrt(float64(n)))
	for i := 2; i <= k; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true

}