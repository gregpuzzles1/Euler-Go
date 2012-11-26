package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	tally := 0
	for y := 1; y <= 10000; y++ {
		x := factors(y)
		x1 := get_total(x)
		x3 := factors(x1)
		x4 := get_total(x3)
		if y == x4 && y != x1 {
			tally += y
		}
	}
	fmt.Println(tally)
}

func get_total(factors []int) (number int) {
	total := 0
	for i := 0; i < len(factors); i++ {
		total += factors[i]
	}
	return total
}

func factors(number int) (factors []int) {
	L := make([]int, number)
	counter := 0
	for i := 1; i < ((number / 2) + 1); i++ {
		if number%i == 0 {
			L[counter] = i
			counter += 1
		}
	}
	return L[:counter]

}
