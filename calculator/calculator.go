package calculator

import "fmt"

func Calculator() {
	//price
	price := []float64{10, 20, 30}
	//tax rates
	tax_rate := []float64{0, 0.07, 0.1, 0.15}
	// result map with tax rate as a key and price included taxrate as value of array
	result := make(map[float64][]float64)

	for _, tax_value := range tax_rate {
		// price included result
		price_included := make([]float64, len(price))
		for price_idx, price_value := range price {
			price_included[price_idx] = price_value * (1 + tax_value)
		}
		// save price included in result as key value pair
		result[tax_value] = price_included
	}

	fmt.Println(result)

}
