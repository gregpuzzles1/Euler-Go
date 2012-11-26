package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")

	L := make(map[int]string)

	L[1] = "Sun"
	L[2] = "Mon"
	L[3] = "Tue"
	L[4] = "Wed"
	L[5] = "Thurs"
	L[6] = "Fri"
	L[7] = "Sat"

	// Start at Monday, 1st, 1900
	counter := 1
	tally := 0

	for yr := 1900; yr < 2001; yr++ {	// Year
		fmt.Printf("Year = %d\n", yr)

		for month := 1; month < 13; month++ {	//  Month
			y := find_days(month, yr)
			for month_days := 1; month_days <= y; month_days++ {	// Number of month days
				// Start at Monday, 1st, 1900
				counter += 1
				if month_days == 1 && L[counter] == "Sun" {
					fmt.Printf("month_days = %d L[counter] = %s\n", month_days, L[counter])
					tally += 1
				}
				if counter >= 7 {	// Days of the week
					counter = 0
				}
			}
		}
		fmt.Printf("\n")
	}
	// subtract 2 for the year 1900
	fmt.Println("Tally = ", tally-2)
}

func find_days(month int, yr int) int {
	d := 0
	if (month == 1) || (month == 3) || (month == 5) || (month == 7) || (month == 8) || (month == 10) || (month == 12) {
		d = 31
	} else if (month == 4) || (month == 6) || (month == 9) || (month == 11) {
		d = 30
	} else if month == 2 {
		if yr%4 == 0 {
			if yr%100 == 0 {
				if yr%400 == 0 {
					d = 29
				} else {
					d = 28
				}
			} else {
				d = 29
			}
		} else {
			d = 28
		}
	} else {
		fmt.Println("There has been an error")
	}
	return d
}

