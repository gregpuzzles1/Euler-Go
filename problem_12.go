package main

import "fmt"
import "math"

func main() {
	counter := 0
	TriangleNumber := 0
	Switch := 1
	for Switch >= 1 {
		counter++
		TriangleNumber += counter
		if factors(TriangleNumber) > 500 {
			fmt.Println(TriangleNumber)
			Switch = 0
		}
	}
}

func factors(n int) (facCount int) {
	facCounter := 0
	k := int(math.Sqrt(float64(n)))
	for i := 1; i < k+1; i++ {
		if n%i == 0 {
			facCounter++
		}
	}
	return facCounter * 2

}
