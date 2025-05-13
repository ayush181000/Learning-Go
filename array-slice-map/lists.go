package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	productNames := [4]string{"A book"}
	prices := [4]float64{10.99, 11.99, 20.00, 0}
	productNames[2] = "A carpet"
	fmt.Println(prices)
	fmt.Println(productNames)
}
