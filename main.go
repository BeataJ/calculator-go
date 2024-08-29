package main

import (
	// "fmt"

	"example.com/calculator/cmdmanager"
	// "example.com/calculator/filemanager"
	"example.com/calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}

}
