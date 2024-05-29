package v6

// Predefined swap modes.
const (
	SwapModeExactIn  = "ExactIn"
	SwapModeExactOut = "ExactOut"
)

const (
	DexCloneProtocol   = "Clone Protocol"
	DexSaros           = "Saros"
	DexSanctum         = "Sanctum"
	DexRaydiumCP       = "Raydium CP"
	DexOpenBookV2      = "OpenBook V2"
	DexCropper         = "Cropper"
	DexRaydiumCLMM     = "Raydium CLMM"
	DexAldrinV2        = "Aldrin V2"
	DexMeteoraDLMM     = "Meteora DLMM"
	DexMarinade        = "Marinade"
	DexPenguin         = "Penguin"
	DexTokenSwap       = "Token Swap"
	DexCrema           = "Crema"
	DexSanctumInfinity = "Sanctum Infinity"
	DexPhoenix         = "Phoenix"
	DexMeteora         = "Meteora"
	DexOrcaV2          = "Orca V2"
	DexAldrin          = "Aldrin"
	DexOrcaV1          = "Orca V1"
	DexPerps           = "Perps"
	DexHeliumNetwork   = "Helium Network"
	DexSaber           = "Saber"
	DexCropperLegacy   = "Cropper Legacy"
	DexRaydium         = "Raydium"
	DexFluxBeam        = "FluxBeam"
	DexGooseFX         = "GooseFX"
	DexOpenbook        = "Openbook"
	DexBonkswap        = "Bonkswap"
	DexSaberDecimals   = "Saber (Decimals)"
	DexInvariant       = "Invariant"
	DexWhirlpool       = "Whirlpool"
	DexStepN           = "StepN"
	DexLifinityV2      = "Lifinity V2"
	DexLifinityV1      = "Lifinity V1"
	DexDexlab          = "Dexlab"
	DexOasis           = "Oasis"
	DexMercuria        = "Mercuria"
)

// QuoteParams are the parameters for a quote request.
type QuoteParams struct {
	InputMint  string `url:"inputMint"`  // required. Input token mint address
	OutputMint string `url:"outputMint"` // required. Output token mint address
	Amount     uint64 `url:"amount"`     // required. The amount to swap, have to factor in the token decimals.

	SlippageBps                   uint64   `url:"slippageBps,omitempty"`                   // Default is 50 unless autoSlippage is set to true. The slippage % in BPS. If the output token amount exceeds the slippage then the swap transaction will fail.
	SwapMode                      string   `url:"swapMode,omitempty"`                      // (ExactIn or ExactOut) Defaults to ExactIn. ExactOut is for supporting use cases where you need an exact token amount, like payments. In this case the slippage is on the input token.
	Dexes                         []string `url:"dexes,omitempty"`                         // Default is that all DEXes are included. You can pass in the DEXes that you want to include only and separate them by. You can check out the full list here. https://quote-api.jup.ag/v6/program-id-to-label
	ExcludeDexes                  []string `url:"excludeDexes,omitempty"`                  // Default is that all DEXes are included. You can pass in the DEXes that you want to exclude and separate them by. You can check out the full list here.
	RestrictIntermediateTokens    bool     `url:"restrictIntermediateTokens,omitempty"`    // Restrict intermediate tokens to a top token set that has stable liquidity. This will help to ease potential high slippage error rate when swapping with minimal impact on pricing.
	OnlyDirectRoutes              bool     `url:"onlyDirectRoutes,omitempty"`              // Default is false. Direct Routes limits Jupiter routing to single hop routes only.
	AsLegacyTransaction           bool     `url:"asLegacyTransaction,omitempty"`           // Default is false. Instead of using versioned transaction, this will use the legacy transaction.
	PlatformFeeBps                uint64   `url:"platformFeeBps,omitempty"`                // If you want to charge the user a fee, you can specify the fee in BPS. Fee % is taken out of the output token.
	maxAccounts                   uint64   `url:"maxAccounts,omitempty"`                   // Rough estimate of the max accounts to be used for the quote, so that you can compose with your own accounts
	AutoSlippage                  bool     `url:"autoSlippage,omitempty"`                  // Default is false. By setting this to true, our API will suggest smart slippage info that you can use. computedAutoSlippage is the computed result, and slippageBps is what we suggest you to use. Additionally, you should check out maxAutoSlippageBps and autoSlippageCollisionUsdValue.
	maxAutoSlippageBps            uint64   `url:"maxAutoSlippageBps,omitempty"`            // In conjunction with autoSlippage=true, the maximum slippageBps returned by the API will respect this value. It is recommended that you set something here.
	autoSlippageCollisionUsdValue uint64   `url:"autoSlippageCollisionUsdValue,omitempty"` //If autoSlippage is set to true, our API will use a default 1000 USD value as way to calculate the slippage impact for the smart slippage. You can set a custom USD value using this parameter.
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
	QuoteResponse *QuoteResponse `json:"quoteResponse"` // required
	UserPublicKey string         `json:"userPublicKey"` // required

	WrapAndUnwrapSol              *bool  `json:"wrapAndUnwrapSol,omitempty"`              // Default is true. If true, will automatically wrap/unwrap SOL. If false, it will use wSOL token account. Will be ignored if destinationTokenAccount is set because the destinationTokenAccount may belong to a different user that we have no authority to close.
	UseSharedAccounts             *bool  `json:"useSharedAccounts,omitempty"`             // Default is true. This enables the usage of shared program accountns. That means no intermediate token accounts or open orders accounts need to be created for the users. But it also means that the likelihood of hot accounts is higher.
	FeeAccount                    string `json:"feeAccount,omitempty"`                    // Fee token account, same as the output token for ExactIn and as the input token for ExactOut, it is derived using the seeds = ["referral_ata", referral_account, mint] and the REFER4ZgmyYx9c6He5XfaTMiGfdLwRnkV4RPp9t9iF3 referral contract (only pass in if you set a feeBps and make sure that the feeAccount has been created).
	TrackingAccount               string `json:"trackingAccount,omitempty"`               // Tracking account, this can be any public key that you can use to track the transactions, especially useful for integrator. Then, you can use the https://stats.jup.ag/tracking-account/:public-key/YYYY-MM-DD/HH endpoint to get all the swap transactions from this public key.
	ComputeUnitPriceMicroLamports *int64 `json:"computeUnitPriceMicroLamports,omitempty"` // The compute unit price to prioritize the transaction, the additional fee will be computeUnitLimit (1400000) * computeUnitPriceMicroLamports. If auto is used, Jupiter will automatically set a priority fee and it will be capped at 5,000,000 lamports / 0.005 SOL.
	PrioritizationFeeLamports     *int64 `json:"prioritizationFeeLamports,omitempty"`     // Prioritization fee lamports paid for the transaction in addition to the signatures fee. Mutually exclusive with compute_unit_price_micro_lamports. If auto is used, Jupiter will automatically set a priority fee and it will be capped at 5,000,000 lamports / 0.005 SOL. If autoMultiplier ({"autoMultiplier"}: 3}) is used, the priority fee will be a multplier on the auto fee. If jitoTipLamports ({"jitoTipLamports": 5000}) is used, a tip intruction will be included to Jito and no priority fee will be set.
	AsLegacyTransaction           *bool  `json:"asLegacyTransaction,omitempty"`           // Default is false. Request a legacy transaction rather than the default versioned transaction, needs to be paired with a quote using asLegacyTransaction otherwise the transaction might be too large.
	UseTokenLedger                *bool  `json:"useTokenLedger,omitempty"`                // Default is false. This is useful when the instruction before the swap has a transfer that increases the input token amount. Then, the swap will just use the difference between the token ledger token amount and post token amount.
	DestinationTokenAccount       string `json:"destinationTokenAccount,omitempty"`       // Public key of the token account that will be used to receive the token out of the swap. If not provided, the user's ATA will be used. If provided, we assume that the token account is already initialized.
	DynamicComputeUnitLimit       *bool  `json:"dynamicComputeUnitLimit,omitempty"`       // When enabled, it will do a swap simulation to get the compute unit used and set it in ComputeBudget's compute unit limit. This will increase latency slightly since there will be one extra RPC call to simulate this. Default is false.
	SkipUserAccountsRpcCalls      *bool  `json:"skipUserAccountsRpcCalls,omitempty"`      // When enabled, it will not do any rpc calls check on user's accounts. Enable it only when you already setup all the accounts needed for the trasaction, like wrapping or unwrapping sol, destination account is already created.
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
