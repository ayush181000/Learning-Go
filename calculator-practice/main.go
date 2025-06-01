package main

import (
	"calculator/filemanager"
	"calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices1.txt", fmt.Sprintf("result-%.0f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Process failed")
			fmt.Println(err)
		}
	}

}
