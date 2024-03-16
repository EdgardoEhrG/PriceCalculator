package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	TaxRate float64
	Prices []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPricesJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices = prices
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		Prices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}