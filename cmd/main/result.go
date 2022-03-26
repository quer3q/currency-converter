package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func printResult(a *args, code int, body []byte) {
	res := &Response{}
	err := json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("can't unmarshall body to struct")
		fmt.Println(fmt.Sprintf("STATUS: %d", code))
		fmt.Println(fmt.Sprintf("BODY: %s", string(body)))
		return
	}

	v, ok := res.Data[a.symbol]
	if !ok {
		fmt.Println(fmt.Sprintf("no data provided for %s", a.symbol))
		return
	}

	amount, ok := v.Quote[a.convert]
	if !ok {
		fmt.Println(fmt.Sprintf("no amount provided for %s", a.convert))
		return
	}

	fmt.Println(strconv.FormatFloat(amount.Price, 'f', 2, 64))
}
