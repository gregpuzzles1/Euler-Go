package main

import "fmt"

var (
	sum int = 0
)

func main() {
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}