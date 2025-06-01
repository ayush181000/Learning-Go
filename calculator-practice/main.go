package main

import (
	"calculator/filemanager"
	"calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for taxRateIndex, taxRate := range taxRates {
		doneChans[taxRateIndex] = make(chan bool)
		errorChans[taxRateIndex] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result-%.0f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[taxRateIndex], errorChans[taxRateIndex])

		// if err != nil {
		// 	fmt.Println("Process failed")
		// 	fmt.Println(err)
		// }
	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			{
				if err != nil {
					fmt.Println("Process failed")
					fmt.Println(err)
				}
			}
		case <-doneChans[index]:
			{
				fmt.Println("Process done")
			}
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }

}
