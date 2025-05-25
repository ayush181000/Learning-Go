package main

import "fmt"

func main() {

	fmt.Println(factorial(5))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(accumulator(0, numbers...))

	fmt.Println(accumulator(50, 1, -1, 2, 4))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func accumulator(startingValue int, restValues ...int) int {
	sum := startingValue
	for _, val := range restValues {
		sum += val
	}

	return sum
}
