package main

import "fmt"
import "math/big"

func main() {

	n := big.NewInt(40)	// The total number of moves for any one path (right + down)
	r := big.NewInt(20)	// The total number of right moves for any one path

	t := big.NewInt(0).Sub(n, r)				// Part of Denominator
	x := big.NewInt(0).Mul(factorial(r), factorial(t))	// The complete Denominator
	Answer := big.NewInt(0).Div(factorial(n), x)		// Divide the Numerator by the Denominator

	fmt.Println(Answer)

}

func factorial(n *big.Int) (result *big.Int) {
	// This function finds the factorial of a number. For example: 5! = 5 * 4 * 3 * 2 * 1 = 120

	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
}
