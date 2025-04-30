package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(filename string) float64 {
	data, _ := os.ReadFile(filename)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func WriteToFile(accountBalanceFile string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(accountBalanceFile, []byte(valueText), 0644)
}
