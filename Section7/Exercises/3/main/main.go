package main

import "fmt"

func main() {
	nums := []int{5, 4, 2, 3, 1}
	fmt.Printf("Greatest from 5, 4, 2, 3, 1 is %v\n", findGreatest(nums...))
	nums = []int{2, 2, 1, -4, 6}
	fmt.Printf("Greatest from 2,2,1,-4,6 is %v\n", findGreatest(nums...))

}

func findGreatest(numberList ...int) int {
	max := 0
	for _, num := range numberList {
		if num > max {
			max = num
		}

	}
	return max

}
