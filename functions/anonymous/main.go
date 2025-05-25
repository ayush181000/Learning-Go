package main

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)

	transformNumbers(&numbers, getTransformerFunction())
	transformNumbers(&numbers, func(number int) int { return number * 2 }) // this is anonymous functions
	transformNumbers(&numbers, createTransformer(2))

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

func createTransformer(factor int) transformFn {
	return func(x int) int {
		return x * factor
	}
}
