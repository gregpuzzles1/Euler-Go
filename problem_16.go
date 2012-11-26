package main

import "fmt"
import "math/big"
import "strings"
import "strconv"

func main() {

	x := big.NewInt(2)
	y := big.NewInt(1000)
	z := big.NewInt(0)

	result := new(big.Int).Exp(x, y, z)		// result is a 302 digit big.Int
	string_result := result.String()		// convert result to a string
	numbers := strings.Split(string_result, "")	// split the string

	Answer := 0
	for _, number := range numbers {	// enumerate thru the string of numbers
		x, _ := strconv.Atoi(number)	// convert string to an integer
		Answer += x			// add the individual integer to the Answer
	}

	fmt.Println(Answer)
}
