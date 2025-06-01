package cmdmanager

import "fmt"

type CmdManager struct{}

func New() CmdManager {
	return CmdManager{}
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with an ENTER key.")
	var prices []string

	for {
		var price string
		fmt.Scanln(&price)
		if price == "" || price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) WriteJson(data any) error {
	fmt.Println(data)
	return nil
}
