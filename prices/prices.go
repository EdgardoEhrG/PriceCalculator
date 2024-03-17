package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	Prices []float64	`json:"prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

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

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: fm,
		Prices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}