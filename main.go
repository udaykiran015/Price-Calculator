package main

import (
	
	"practice/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.15}

	for _, taxrate := range taxRates {
		priceJob :=prices.NewTaxIncludePriceJob(taxrate)
		priceJob.Process()	}
	

}
