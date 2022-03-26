package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	priceConversionURL = "https://sandbox-api.coinmarketcap.com/v2/tools/price-conversion"
	sandboxAPIKey      = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
)

func sendRequest(a *args) (int, []byte, error) {
	req, err := http.NewRequest("GET", priceConversionURL, nil)
	if err != nil {
		return 0, nil, err
	}

	q := url.Values{}
	q.Add("amount", a.amount)
	q.Add("symbol", a.symbol)
	q.Add("convert", a.convert)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", sandboxAPIKey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err, "can't send request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "can't parse response body")
	}

	return resp.StatusCode, body, nil
}
