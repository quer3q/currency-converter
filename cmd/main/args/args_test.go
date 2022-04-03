package args

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseArgs(t *testing.T) {
	t.Parallel()

	t.Run("error for wrong args length", func(t *testing.T) {
		t.Parallel()
		res, err := ParseArgs([]string{"127.12", "BTC"})

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("error for wrong float", func(t *testing.T) {
		t.Parallel()
		res, err := ParseArgs([]string{"no-float", "BTC", "USD"})

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("no errors for valid values", func(t *testing.T) {
		t.Parallel()
		amount := "123.45"
		symbol := "BTC"
		convert := "USD"

		res, err := ParseArgs([]string{amount, symbol, convert})

		require.NoError(t, err)
		require.Equal(t, amount, res.Amount)
		require.Equal(t, symbol, res.Symbol)
		require.Equal(t, convert, res.Convert)
	})
}
