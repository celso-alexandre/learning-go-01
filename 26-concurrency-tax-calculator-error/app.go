package main

import (
	"fmt"

	"github.com/celso-alexandre/learning-go-01/26-concurrency-tax-calculator-error/filemanager"
	"github.com/celso-alexandre/learning-go-01/26-concurrency-tax-calculator-error/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select { // Select is a switch-case to identify which of 2 or more channels emitted a value first (Promise.race in JavaScript)
		case <-doneChans[i]:
			// Do nothing
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
