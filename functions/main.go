package main

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)

	transformNumbers(&numbers, getTransformerFunction())

	fmt.Println(numbers)

}

func transformNumbers(numbers *[]int, transform transformFn) {
	for i := range *numbers {
		(*numbers)[i] = transform((*numbers)[i])
	}
}

func getTransformerFunction() transformFn {
	return square
}

func square(x int) int {
	return x * x
}
