package main

import "fmt"
import "math"

func main() {
	for n := 1; n < 500; n++ {
		for m := (n + 1); m < 500; m++ {
			a := int(math.Pow(float64(m), 2) - math.Pow(float64(n), 2))
			b := 2 * m * n
			c := int(math.Pow(float64(m), 2) + math.Pow(float64(n), 2))
			if a+b+c == 1000 {
				product := a * b * c
				fmt.Println(a, b, c)
				fmt.Println(product)
			}
		}

	}
}