package main

import (
	"fmt"
)

func evenlyDivisible(n int) (eD int) {
	/*Checks if number (n) is evenly divisible by the numbers 1 -20.*/

	counter := 0
	for i := 1; i < 21; i++ {
		if n%i == 0 {
			counter += 1
			if counter == 20 {
				return n
			}

		} else {
			return 1
		}
	}
	return 1
}
func main() {
	/*Main program - since we know that 2520 is evenly divisible by
	the first 10 integers, we know that the number in question is a
	multiple of 2520. So we start with 2520 and increment by that much
	until we find the smallest number that is evenly divisible by the
	first 20 integers (1-20)*/

	n := 2520
	Switch := 1
	for Switch >= 1 {
		smallest := evenlyDivisible(n)
		if smallest != 1 {
			fmt.Println(smallest)
			Switch = 0

		} else {
			n += 2520
		}
	}
}
