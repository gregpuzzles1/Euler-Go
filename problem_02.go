package main

import "fmt"

func fib(n int) int {
	thesum := 0
	a := 1
	b := 1
	for a < n {
		if a%2 == 0 {
			thesum = thesum + a
		}
		a, b = b, a+b

	}
	return thesum

}

func main() {
	fmt.Println(fib(4000000))
}