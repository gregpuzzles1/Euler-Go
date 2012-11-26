package main

import "fmt"
import "strconv"
import "strings"

func isPalindrome(product int) (isP bool) {
	s := strconv.Itoa(product)
	a := strings.Split(s, "")
	L := len(a)
	for i := 0; i < L/2; i++ {
		a[i], a[L-i-1] = a[L-i-1], a[i]
	}
	x := strings.Join(a, "")
	y, _ := strconv.Atoi(x)
	if product == y {
		return true
	}
	return false

}

func main() {
	largest := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) {
				if product > largest {
					largest = product
				}

			}
		}

	}
	fmt.Println(largest)
}