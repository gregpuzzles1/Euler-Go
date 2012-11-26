package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("Hello, playground")

	number := make(map[int]string)

	number[1] = "one"
	number[2] = "two"
	number[3] = "three"
	number[4] = "four"
	number[5] = "five"
	number[6] = "six"
	number[7] = "seven"
	number[8] = "eight"
	number[9] = "nine"
	number[10] = "ten"
	number[11] = "eleven"
	number[12] = "twelve"
	number[13] = "thirteen"
	number[14] = "fourteen"
	number[15] = "fifteen"
	number[16] = "sixteen"
	number[17] = "seventeen"
	number[18] = "eighteen"
	number[19] = "nineteen"
	number[20] = "twenty"
	number[30] = "thirty"
	number[40] = "forty"
	number[50] = "fifty"
	number[60] = "sixty"
	number[70] = "seventy"
	number[80] = "eighty"
	number[90] = "ninety"
	number[100] = "hundred"
	number[1000] = "thousand"

	tally := 0
	k := 0

	for n := 1; n <= 1000; n++ {
		x := strconv.FormatInt(int64(n), 10)
		if len(x) == 1 {
			// The first 9 numbers are 1 digit -- in the dictionary
			k = len(number[n])
			fmt.Printf("n = %d k = %d\n", n, k)
		} else if len(x) == 2 {
			x1 := x[0:1]
			x2 := x[1:2]
			x2_i, _ := strconv.Atoi(x2)
			if n < 20 {
				k = len(number[n])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else {
				if x2 == "0" {
					// The numbers under 100 and greater than 19 -- in the 
					// dictionary ending in '0', (20, 30, 40 ....)
					k = len(number[n])
					fmt.Printf("n = %d k = %d\n", n, k)
				} else {
					// The other numbers under 100 greater than 19
					x1a := x1 + "0"
					x1a_i, _ := strconv.Atoi(x1a)
					k = len(number[x1a_i]) + len(number[x2_i])
					fmt.Printf("n = %d k = %d\n", n, k)
				}
			}
		} else if len(x) == 3 {
			// add 3 for 'and' i.e. -- two-hundred and ten
			x1 := x[0:1]
			x2 := x[1:2]
			x3 := x[2:3]
			x1a := x2 + "0"
			x1aa := x2 + x3

			x1_i, _ := strconv.Atoi(x1)
			x3_i, _ := strconv.Atoi(x3)
			x1a_i, _ := strconv.Atoi(x1a)
			x1aa_i, _ := strconv.Atoi(x1aa)

			if x1 == "1" && x2 == "0" && x3 == "0" {
				// 100 -- in the dictionary
				k = len(number[1]) + len(number[100])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else if x2 == "0" && x3 == "0" {
				// Consider 200, 300, 400, 500, 600, 700, 800, and 900
				k = len(number[x1_i]) + len(number[100])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else if x2 == "0" && x3 != "0" {
				// Consider 101, 102 ... 109, 201, 202, ... 209 etc.
				k = 3 + len(number[x1_i]) + len(number[100]) + len(number[x3_i])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else if x2 != "0" && x3 == "0" {
				// Consider 110, 120, ... 190, 210, 220, ... 290 etc.
				k = 3 + len(number[x1_i]) + len(number[100]) + len(number[x1a_i])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else if x2 == "1" && x3 != "0" {
				// Consider the teens 111, 112, ... 119, 211, 212, ... 219 etc.
				k = 3 + len(number[x1_i]) + len(number[100]) + len(number[x1aa_i])
				fmt.Printf("n = %d k = %d\n", n, k)
			} else {
				// Consider all the other numbers
				k = 3 + len(number[x1_i]) + len(number[100]) + len(number[x1a_i]) + len(number[x3_i])
				fmt.Printf("n = %d k = %d\n", n, k)
			}

		} else {
			// 1000 -- two parts (one and thousand) -- in the dictionary
			k = len(number[1]) + len(number[1000])
			fmt.Printf("n = %d k = %d\n", n, k)
		}
		tally += k

	}
	fmt.Println(tally)
}
