package main

import "miqgo.com/go/cmasters/api"

func main() {

	rate, error := api.GetRate("BTC")
	if error != nil {
		panic(error)
	}

	print(rate)

}
