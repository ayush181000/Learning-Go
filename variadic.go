package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {

	total := 0
	for i := range nums {
		total = total + nums[i]
	}

	return total
}
