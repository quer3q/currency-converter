package main

import (
	"flag"
	"fmt"

	"github.com/quer3q/currency-converter/cmd/main/coinmarket"

	"github.com/quer3q/currency-converter/cmd/main/args"
)

func main() {
	flag.Parse()

	a, err := args.ParseArgs(flag.Args())
	if err != nil {
		fmt.Println(fmt.Sprintf("can't parse args: %s", err.Error()))
		args.PrintHelp()
		return
	}

	coinmarketService := coinmarket.New()
	res, err := coinmarketService.ConvertPrice(a)
	if err != nil {
		fmt.Println(fmt.Sprintf("an error occured: %s", err.Error()))
	}

	fmt.Println(res)
}
