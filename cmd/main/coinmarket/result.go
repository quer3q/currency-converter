package coinmarket

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"

	"github.com/quer3q/currency-converter/cmd/main/args"
	"github.com/quer3q/currency-converter/pkg"
)

func checkResult(a *args.Args, body []byte) (string, error) {
	res := &pkg.Response{}
	err := json.Unmarshal(body, res)
	if err != nil {
		return "", errors.Wrap(err, "can't unmarshall body to struct")
	}

	v, ok := res.Data[a.Symbol]
	if !ok {
		return "", errors.Errorf("no data provided for %s", a.Symbol)
	}

	amount, ok := v.Quote[a.Convert]
	if !ok {
		return "", errors.Errorf("no amount provided for %s", a.Convert)
	}

	return strconv.FormatFloat(amount.Price, 'f', 2, 64), nil
}
