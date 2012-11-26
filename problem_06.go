package main

import "fmt"
import "math"

var (
	sum_of_squares	float64	= 0
	square_of_sum	float64	= 0
	Sum		float64	= 0
)

func main() {
	for i := 1; i < 101; i++ {
		square := math.Pow(float64(i), 2)
		sum_of_squares += square
		Sum += float64(i)
	}
	square_of_sum = math.Pow(float64(Sum), 2)
	fmt.Println("sum_of_squares = ", int(sum_of_squares))
	fmt.Println("square_of_sum = ", int(square_of_sum))
	Difference := square_of_sum - sum_of_squares
	fmt.Println("Difference = ", int(Difference))

}