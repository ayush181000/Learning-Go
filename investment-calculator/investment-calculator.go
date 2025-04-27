package main

import (
	"fmt"
	"math"
)

const inflationRate = 6

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64 = 1.5

	fmt.Print("Invested Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formatSV := fmt.Sprintf("Future Value: %.1f \n", futureValue)
	formatRSV := fmt.Sprintf("Future Value (Adjusted for inflation): %.1f \n", futureRealValue)

	fmt.Print(formatSV, formatRSV)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, futureRealValue float64) {

	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue = futureValue / math.Pow((1+inflationRate/100), years)

	return futureValue, futureRealValue
}
