package main

import "fmt"
import "io/ioutil"
import "strings"
import "sort"

func main() {
	fData, err := ioutil.ReadFile("names.txt")	// read in the external file
	if err != nil {
		fmt.Println("Err is ", err)	// print any error
	}
	strbuffer := string(fData)	// convert read in file to a string

	names := strings.Split(strbuffer, ",")

	for i, val := range names {
		names[i] = strings.Trim(val, "\"")
	}

	sort.Strings(names)

	total := 0

	for i := range names {
		name := names[i]
		letters := strings.Split(name, "")

		x2 := 0
		y := 0

		for m := range letters {
			letter := letters[m]
			x1 := letter_counter(letter)
			x2 += x1
		}
		y = x2 * (i + 1)
		total += y

	}
	fmt.Println(total)

}

func letter_counter(letter string) (lc int) {
	letv := make(map[string]int)

	letv["A"] = 1
	letv["B"] = 2
	letv["C"] = 3
	letv["D"] = 4
	letv["E"] = 5
	letv["F"] = 6
	letv["G"] = 7
	letv["H"] = 8
	letv["I"] = 9
	letv["J"] = 10
	letv["K"] = 11
	letv["L"] = 12
	letv["M"] = 13
	letv["N"] = 14
	letv["O"] = 15
	letv["P"] = 16
	letv["Q"] = 17
	letv["R"] = 18
	letv["S"] = 19
	letv["T"] = 20
	letv["U"] = 21
	letv["V"] = 22
	letv["W"] = 23
	letv["X"] = 24
	letv["Y"] = 25
	letv["Z"] = 26

	lc = letv[letter]

	return lc
}
