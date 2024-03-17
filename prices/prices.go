package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager iomanager.IOManager`json:"-"`
	TaxRate float64 `json:"tax_rate"`
	Prices []float64	`json:"prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.Prices = prices

	return nil
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: iom,
		Prices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}