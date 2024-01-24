package prices

type TaxIncludedPrices struct {
	TaxRate float64
	Prices []float64
	TaxIncludedPrices map[string]float64
}