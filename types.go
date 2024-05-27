package jupiter

// Predefined swap modes.
const (
	SwapModeExactIn  = "ExactIn"
	SwapModeExactOut = "ExactOut"
)

// QuoteParams are the parameters for a quote request.
type QuoteParams struct {
	InputMint  string `url:"inputMint"`  // required
	OutputMint string `url:"outputMint"` // required
	Amount     uint64 `url:"amount"`     // required

	SwapMode            string `url:"swapMode,omitempty"` // Swap mode, default is ExactIn; Available values : ExactIn, ExactOut.
	SlippageBps         uint64 `url:"slippageBps,omitempty"`
	FeeBps              uint64 `url:"feeBps,omitempty"`              // Fee BPS (only pass in if you want to charge a fee on this swap)
	OnlyDirectRoutes    bool   `url:"onlyDirectRoutes,omitempty"`    // Only return direct routes (no hoppings and split trade)
	AsLegacyTransaction bool   `url:"asLegacyTransaction,omitempty"` // Only return routes that can be done in a single legacy transaction. (Routes might be limited)
	UserPublicKey       string `url:"userPublicKey,omitempty"`       // Public key of the user (only pass in if you want deposit and fee being returned, might slow down query)
}

// QuoteResponse is the response from a quote request.
type QuoteResponse struct {
	InputMint            string      `json:"inputMint"`
	InAmount             string      `json:"inAmount"`
	OutputMint           string      `json:"outputMint"`
	OutAmount            string      `json:"outAmount"`
	OtherAmountThreshold string      `json:"otherAmountThreshold"`
	SwapMode             string      `json:"swapMode"`
	SlippageBps          int         `json:"slippageBps"`
	PlatformFee          interface{} `json:"platformFee"`
	PriceImpactPct       string      `json:"priceImpactPct"`
	RoutePlan            []struct {
		SwapInfo struct {
			AmmKey     string `json:"ammKey"`
			Label      string `json:"label"`
			InputMint  string `json:"inputMint"`
			OutputMint string `json:"outputMint"`
			InAmount   string `json:"inAmount"`
			OutAmount  string `json:"outAmount"`
			FeeAmount  string `json:"feeAmount"`
			FeeMint    string `json:"feeMint"`
		} `json:"swapInfo"`
		Percent int `json:"percent"`
	} `json:"routePlan"`
	ContextSlot int     `json:"contextSlot"`
	TimeTaken   float64 `json:"timeTaken"`
}

// SwapParams are the parameters for a swap request.
type SwapParams struct {
	QuoteResponse                 *QuoteResponse `json:"quoteResponse"`                           // required
	UserPublicKey                 string         `json:"userPublicKey"`                           // required
	WrapUnwrapSol                 *bool          `json:"wrapUnwrapSOL,omitempty"`                 // Default is true. If true, will automatically wrap/unwrap SOL. If false, it will use wSOL token account. Will be ignored if destinationTokenAccount is set because the destinationTokenAccount may belong to a different user that we have no authority to close.
	FeeAccount                    string         `json:"feeAccount,omitempty"`                    // Fee token account for the platform fee (only pass in if you set a feeBps), the mint is outputMint for the default swapMode.ExactOut and inputMint for swapMode.ExactIn.
	AsLegacyTransaction           *bool          `json:"asLegacyTransaction,omitempty"`           // Request a legacy transaction rather than the default versioned transaction, needs to be paired with a quote using asLegacyTransaction otherwise the transaction might be too large.
	ComputeUnitPriceMicroLamports *int64         `json:"computeUnitPriceMicroLamports,omitempty"` // Compute unit price to prioritize the transaction, the additional fee will be compute unit consumed * computeUnitPriceMicroLamports.
}

// SwapResponse is the response from a swap request.
type SwapResponse struct {
	SwapTransaction           string `json:"swapTransaction"` // base64 encoded transaction string
	LastValidBlockHeight      int64  `json:"lastValidBlockHeight"`
	PrioritizationFeeLamports int64  `json:"prioritizationFeeLamports"`
}

type SwapInstructionsResp struct {
	TokenLedgerInstruction      interface{}   `json:"tokenLedgerInstruction"`
	ComputeBudgetInstructions   []Instruction `json:"computeBudgetInstructions"`
	SetupInstructions           []Instruction `json:"setupInstructions"`
	SwapInstruction             Instruction   `json:"swapInstruction"`
	CleanupInstruction          Instruction   `json:"cleanupInstruction"`
	OtherInstructions           []Instruction `json:"otherInstructions"`
	AddressLookupTableAddresses []string      `json:"addressLookupTableAddresses"`
	PrioritizationFeeLamports   int64         `json:"prioritizationFeeLamports"`
}
type Instruction struct {
	ProgramId string    `json:"programId"`
	Accounts  []Account `json:"accounts"`
	Data      string    `json:"data"`
}

type Account struct {
	Pubkey     string `json:"pubkey"`
	IsSigner   bool   `json:"isSigner"`
	IsWritable bool   `json:"isWritable"`
}
