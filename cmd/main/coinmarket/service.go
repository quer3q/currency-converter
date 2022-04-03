package coinmarket

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/quer3q/currency-converter/cmd/main/args"
)

type Service interface {
	ConvertPrice(ar *args.Args) (string, error)
}

type impl struct{}

func New() Service {
	return &impl{}
}

func (i *impl) ConvertPrice(ar *args.Args) (string, error) {
	code, body, err := sendRequest(ar)
	if err != nil {
		return "", errors.Wrap(err, "ConvertPrice")
	}
	if code != http.StatusOK {
		return "", errors.Errorf("request failed: %d", code)
	}

	return checkResult(ar, body)
}
