package main

import "fmt"

var a []int
x = a[1]

func sum(a []int) int {   // returns an int
    s := 0;
    for i := 0; i < len(a); i++ {
        s += a[i]
    }
    return s
}

func main() {
	s := sum(3[] ) // a slice of the array is passed to sum
	fmt.Println(s)
	
}


//s := sum(&[3]int{1,2,3}); // a slice of the array is passed to sum