package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

// The flow is mostly synchronous in GO, the only way to cause asynchronousness is to use go
func main() {
	// go getCurrencyData("BTC")
	// go getCurrencyData("ETH")
	// go getCurrencyData("BCH")

	// time.Sleep(time.Second * 5)
	// for {}


	currencies := []string { "BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		// we add one to the waiting group
		// to indicate we have a goroutine running
		wg.Add(1)

		// wg.Done should ideally be called immediately after the go routine has ended
		// in order to do that we can make use of an IIFE to call the getCurrencyData
		// this IIFE can then be made the goroutine, as shown below
		// then remember we pass the currency as an argument to the IIFE, because
		// when the IIFE runs, the for loop has already ended and the currency gets stuck at the last array index
		// in this case "BCH", hence we have to pass it into the IIFE as well
		go func (currencyCode string) {
			getCurrencyData(currencyCode)

			// we call done when the code in the goroutine has finished executing
			wg.Done()
		}(currency)
	}

	// then here we wait for the counter to reach 0
	wg.Wait()

// Note that if the main func ends  all the code ends 
// and all goroutines would have executed without seeing their results
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	// accessing a value from a pointer actually gets the data and returns the value
	// this is for convenience purposes
	// so you don't have to use *rate
	if err == nil {
		fmt.Printf("The rate for currency %v is %.2f \n", rate.Currency, rate.Price)
	}
}