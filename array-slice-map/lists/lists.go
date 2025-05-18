package main

import "fmt"

func main() {
	// productNames := [4]string{"A book"}
	// prices := [4]float64{10.99, 11.99, 20.00, 0}
	// productNames[2] = "A carpet"
	// fmt.Println(prices)
	// fmt.Println(productNames)

	// featuredPrices := prices[1:]
	// highlightedPrices := featuredPrices[:1]
	// featuredPrices[0] = 199.99

	// fmt.Println(featuredPrices)
	// fmt.Println(highlightedPrices)

	// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// highlightedPrices = highlightedPrices[:3]
	// fmt.Println(highlightedPrices)

	prices := []float64{0.1, 0.2}
	fmt.Println(prices[0:1])
	prices[1] = 1.2

	prices = append(prices, 0.3)
	fmt.Println(prices)

	discountedPrices := []float64{101.99, 80.99}

	prices = append(prices, discountedPrices...)
	fmt.Println(prices)

}
