package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Server error")
	}
}
