package main

import (
	
	"practice/prices"
)

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0.1, 0.2, 0.15}
	result := make(map[float64][]float64)

	for _, taxrate := range taxRates {
		pricejob := prices.NewTaxIncludePriceJob(taxrate).Process()
	}
	

}
