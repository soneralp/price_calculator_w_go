package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager.go"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("tax_included_prices_%v.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		priceJob.Process()
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}
	}
}
