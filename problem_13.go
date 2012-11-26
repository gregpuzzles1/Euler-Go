package main

import "fmt"
import "io/ioutil"
import "strings"
import "math/big"

func main() {
	fData, err := ioutil.ReadFile("one-hundred_50.txt")	// read in the external file
	if err != nil {
		fmt.Println("Err is ", err) 	// print any error
	}
	strbuffer := string(fData)	// convert read in file to a string
	lines := strings.Split(strbuffer, "\n") // split the string into lines
	tally := big.NewInt(0)		// declare and initialise big number tally
	for _, line := range lines {	// loop through each line in lines - 100 of them
		bi := big.NewInt(0)	// declare and  initialise big number bi
		bi.SetString(line, 10) 	// convert line to a big integer (base 10)
		tally.Add(tally, bi)   	// add bi to the tally
	}
	fmt.Println(tally) 	// print tally and use the first 10 digits for the answer
}

