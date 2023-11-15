package main

import (
	"fmt"

	"miqgo.com/go/cmasters/api"
)

func main() {

	rate, error := api.GetRate("BTC")
	if error != nil {
		panic(error)
	}

	fmt.Printf("1 BTC = %f USD\n", rate.Price)

}
