package jupiter_test

import (
	"github.com/qiruos/jupiter/utils"
	"testing"

	"github.com/qiruos/jupiter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	wSolMint = "So11111111111111111111111111111111111111112"
	usdcMint = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
)

func TestQuote(t *testing.T) {
	c := jupiter.NewClient()
	quotes, err := c.Quote(jupiter.QuoteParams{
		InputMint:        wSolMint,
		OutputMint:       usdcMint,
		Amount:           100000,
		OnlyDirectRoutes: true,
		SwapMode:         jupiter.SwapModeExactIn,
	})
	require.NoError(t, err)
	require.NotEmpty(t, quotes)

	//utils.PrettyPrint(quotes)

	assert.Equal(t, wSolMint, quotes.InputMint)
	assert.Equal(t, usdcMint, quotes.OutputMint)
	assert.Equal(t, "100000", quotes.InAmount)
}

func TestSwap(t *testing.T) {
	c := jupiter.NewClient()
	var quoteResponse *jupiter.QuoteResponse
	var err error

	t.Run("get quotes", func(t *testing.T) {
		quoteResponse, err = c.Quote(jupiter.QuoteParams{
			InputMint:        wSolMint,
			OutputMint:       usdcMint,
			Amount:           100000,
			OnlyDirectRoutes: false,
		})
		require.NoError(t, err)
		require.NotEmpty(t, quoteResponse)
	})

	t.Run("create swap tx", func(t *testing.T) {
		swapTx, err := c.Swap(jupiter.SwapParams{
			UserPublicKey: "8HwPMNxtFDrvxXn1fJsAYB258TnA6Ydr1DWCtVYgRW4W",
			QuoteResponse: quoteResponse,
			WrapUnwrapSol: utils.Pointer(true),
		})
		require.NoError(t, err)
		require.NotEmpty(t, swapTx)

		//t.Log(swapTx)
	})
}

func TestSwapInstructions(t *testing.T) {
	c := jupiter.NewClient()
	var quoteResponse *jupiter.QuoteResponse
	var err error

	t.Run("get quotes", func(t *testing.T) {
		quoteResponse, err = c.Quote(jupiter.QuoteParams{
			InputMint:        wSolMint,
			OutputMint:       usdcMint,
			Amount:           100000,
			OnlyDirectRoutes: false,
		})
		require.NoError(t, err)
		require.NotEmpty(t, quoteResponse)
	})

	t.Run("create swap tx", func(t *testing.T) {
		swapInstructions, err := c.SwapInstructions(jupiter.SwapParams{
			UserPublicKey: "8HwPMNxtFDrvxXn1fJsAYB258TnA6Ydr1DWCtVYgRW4W",
			QuoteResponse: quoteResponse,
			WrapUnwrapSol: utils.Pointer(true),
		})
		require.NoError(t, err)
		require.NotEmpty(t, swapInstructions)

		//t.Log(swapInstructions)
	})
}
