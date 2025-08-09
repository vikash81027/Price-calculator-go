package main

import (
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxVal := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxVal*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxVal)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
