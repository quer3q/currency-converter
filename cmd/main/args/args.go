package args

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

const (
	requiredArgsLength = 3 // like '123.45 USD BTC'
)

type Args struct {
	Amount  string
	Symbol  string
	Convert string
}

func ParseArgs(in []string) (*Args, error) {
	if len(in) < requiredArgsLength {
		return nil, errors.Errorf("invalid arguments length: %d", len(in))
	}

	amount := in[0]
	// Check that we've got float in args
	if _, err := strconv.ParseFloat(amount, 10); err != nil {
		return nil, errors.Wrapf(err, "can't parse amount to float: %s", amount)
	}

	return &Args{
		Amount:  amount,
		Symbol:  in[1],
		Convert: in[2],
	}, nil
}

func PrintHelp() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s\n"+
		"Example:\n"+
		"./currency-converter 123.45 USD BTC\n"+
		"You provide amount (1) in currency (2) to which convert (3)\n", os.Args[0])
}
