package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"miqgo.com/go/cmasters/datatypes"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiURL, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var cryptoRate datatypes.Rate
		err = json.Unmarshal(bodyBytes, &cryptoRate)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("API returned status code %d", res.StatusCode)
	}

	rate := datatypes.Rate{
		Currency: currency,
		Price:    0.0,
	}

	return &rate, nil

}
