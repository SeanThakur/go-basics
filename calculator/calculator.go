package calculator

import (
	"seanThakur.com/app/calculator/prices"
)

func Calculator() {
	//tax rates
	tax_rate := []float64{0, 0.07, 0.1, 0.15}

	for _, tax_value := range tax_rate {
		// price included result
		job := prices.NewTaxIncludedPriceJob(tax_value)
		job.Process()
	}

}
