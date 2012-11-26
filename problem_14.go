package main

import "fmt"

func main() {
	largest := 0
	longest_sequence := 0
	for n := 2; n < 1000000; n++ {

		var (
			Number int64 = 0
		)
		Number = int64(n)
		counter := 1
		for Number > 1 {
			if Number%2 == 0 {
				Number = Number / 2
			} else {
				Number = Number*3 + 1
			}
			counter += 1
		}

		if counter > longest_sequence {
			longest_sequence = counter
			largest = n
		}

	}
	fmt.Println("Answer = ", largest)
	fmt.Println("longest sequence = ", longest_sequence)

}
