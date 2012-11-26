package main

import "fmt"
import "math/big"
import "strings"
import "strconv"

func main() {

	n := big.NewInt(100)

	result := factorial(n)
	string_result := result.String()		// convert result to a string
	numbers := strings.Split(string_result, "")	// split the string

	Answer := 0
	for _, number := range numbers {	// enumerate thru the string of numbers
		x, _ := strconv.Atoi(number)	// convert string to an integer
		Answer += x			// add the individual integer to the Answer
	}

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
