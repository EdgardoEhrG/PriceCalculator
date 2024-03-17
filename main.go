package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	"price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// ----- To use file with prices

		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))

		// ----- To use terminal and to scan prices

		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}