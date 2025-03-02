package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println(a)

	for i := range len(a) {
		fmt.Println(a[i])
	}

	b := []int{5, 4, 3, 2, 1}

	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
