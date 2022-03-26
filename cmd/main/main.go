package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	a, err := parseArgs(flag.Args())
	if err != nil {
		fmt.Println(fmt.Sprintf("can't parse args: %s", err.Error()))
		printHelp()
		return
	}

	statusCode, resp, err := sendRequest(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	printResult(a, statusCode, resp)
}
