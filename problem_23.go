package main

import "fmt"

func main() {
	fmt.Println("Greg is cool!")
	fmt.Println(sum_all_integers())
	//fmt.Println(sum(abundant_numbers()))	
	//fmt.Println(proper_divisors(28))
	//fmt.Println(sum(proper_divisors(28)))
	//fmt.Println(sum_all_integers() - sum(abundant_numbers()))
	fmt.Println(sum(abundant_sums(abundant_numbers())))
}


func abundant_sums(aN []int) []int{
	counter := 0
	zzz := 0
	aS := make([]int, 1000000)
	for i := range aN {
	//for i := 0; i <= len(aN); i ++ {
		for j := range aN {
		//for j := 0; j <= len(aN); j ++ {
			fmt.Println("aN[i] = ", aN[i])
			fmt.Println("aN[j] = ", aN[j])
			zzz = aN[i] + aN[j]
			if zzz < 28123 {
				aS[counter] = zzz
				counter += 1
			}
		}
	}
	return aS[:counter]
}

func abundant_numbers() []int{
	number := 281
	counter := 1
	L := make([]int, number)
	for i := 1; i <= number; i ++{
		if sum(proper_divisors(i)) > i{
			L[counter] = i
			counter += 1
		}
	}
	return L[:counter]
}

func sum(pd []int) (int){
	sum := 0
	for i := range pd {
		sum += pd[i]
	}
	return sum
}

func proper_divisors(number int) (factors []int) {
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

func sum_all_integers() int {
	sai := 0
	for i := 1; i < 28123; i ++ {
		sai += i
	}
	return sai
}
