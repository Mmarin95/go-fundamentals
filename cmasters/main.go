package main

import (
	"fmt"
	"sync"

	"miqgo.com/go/cmasters/api"
)

func main() {

	currencies := []string{"BTC", "ETH", "BHC"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(curr string) {
			getCurrencyFor(curr)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyFor(currency string) {
	rate, error := api.GetRate(currency)
	if error != nil {
		panic(error)
	}

	fmt.Printf("1 %s = %f USD\n", currency, rate.Price)
}
