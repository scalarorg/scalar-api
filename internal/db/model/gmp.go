package model

type CreatedAtDocument struct {
	Week    int64 `bson:"week"`
	Hour    int64 `bson:"hour"`
	Month   int64 `bson:"month"`
	Year    int64 `bson:"year"`
	Ms      int64 `bson:"ms"`
	Day     int64 `bson:"day"`
	Quarter int64 `bson:"quarter"`
}

type LogDocument struct {
	BlockHash        string   `bson:"blockHash"`
	Address          string   `bson:"address"`
	LogIndex         uint     `bson:"logIndex"`
	Data             string   `bson:"data"`
	Removed          bool     `bson:"removed"`
	Topics           []string `bson:"topics"`
	BlockNumber      uint64   `bson:"blockNumber"`
	TransactionIndex uint     `bson:"transactionIndex"`
	TransactionHash  string   `bson:"transactionHash"`
}

type ReceiptDocument struct {
	BlockHash         string        `bson:"blockHash"`
	ContractAddress   *string       `bson:"contractAddress"`
	TransactionIndex  uint          `bson:"transactionIndex"`
	Type              uint8         `bson:"type"`
	Confirmations     uint          `bson:"confirmations"`
	TransactionHash   string        `bson:"transactionHash"`
	GasUsed           string        `bson:"gasUsed"`
	BlockNumber       uint64        `bson:"blockNumber"`
	CumulativeGasUsed string        `bson:"cumulativeGasUsed"`
	From              string        `bson:"from"`
	To                string        `bson:"to"`
	EffectiveGasPrice string        `bson:"effectiveGasPrice"`
	Logs              []LogDocument `bson:"logs"`
	Status            uint8         `bson:"status"`
}

type ReturnValuesDocument struct {
	Sender                     string `bson:"sender"`
	DestinationChain           string `bson:"destinationChain"`
	DestinationContractAddress string `bson:"destinationContractAddress"`
	PayloadHash                string `bson:"payloadHash"`
	Payload                    string `bson:"payload"`
}

type TransactionDocument struct {
	BlockHash            string `bson:"blockHash"`
	YParity              string `bson:"yParity"`
	TransactionIndex     uint   `bson:"transactionIndex"`
	Type                 uint8  `bson:"type"`
	Nonce                uint64 `bson:"nonce"`
	R                    string `bson:"r"`
	S                    string `bson:"s"`
	ChainID              int64  `bson:"chainId"`
	V                    uint64 `bson:"v"`
	BlockNumber          uint64 `bson:"blockNumber"`
	Gas                  string `bson:"gas"`
	MaxPriorityFeePerGas string `bson:"maxPriorityFeePerGas"`
	From                 string `bson:"from"`
	To                   string `bson:"to"`
	MaxFeePerGas         string `bson:"maxFeePerGas"`
	Value                uint64 `bson:"value"`
	Hash                 string `bson:"hash"`
	GasPrice             string `bson:"gasPrice"`
}

type GMPStepDocument struct {
	Chain                string               `bson:"chain"`
	ContractAddress      string               `bson:"contract_address"`
	Address              string               `bson:"address"`
	Topics               []string             `bson:"topics"`
	BlockNumber          uint64               `bson:"blockNumber"`
	TransactionHash      string               `bson:"transactionHash"`
	TransactionIndex     uint                 `bson:"transactionIndex"`
	BlockHash            string               `bson:"blockHash"`
	LogIndex             uint                 `bson:"logIndex"`
	Removed              bool                 `bson:"removed"`
	ID                   string               `bson:"id"`
	Event                string               `bson:"event"`
	EventSignature       string               `bson:"eventSignature"`
	ReturnValues         ReturnValuesDocument `bson:"returnValues"`
	ChainType            string               `bson:"chain_type"`
	DestinationChainType string               `bson:"destination_chain_type"`
	CreatedAt            CreatedAtDocument    `bson:"created_at"`
	EventIndex           uint                 `bson:"eventIndex"`
	BlockTimestamp       int64                `bson:"block_timestamp"`
	Receipt              ReceiptDocument      `bson:"receipt"`
	Transaction          TransactionDocument  `bson:"transaction"`
	LogIndexAlias        uint                 `bson:"_logIndex"`
	TypeAlias            string               `bson:"_type"`
}

type ConfirmDocument struct {
	BlockNumber           uint64 `bson:"blockNumber"`
	BlockTimestamp        int64  `bson:"block_timestamp"`
	ConfirmaionTxHash     string `bson:"confirmation_txhash"`
	Event                 string `bson:"event"`
	PollId                string `bson:"poll_id"`
	SourceChain           string `bson:"sourceChain"`
	SourceTransactionHash string `bson:"sourceTransactionHash"`
	TransactionHash       string `bson:"transactionHash"`
	TransactionIndex      uint64 `bson:"transactionIndex"`
}

type GasPriceInUnitsDocument struct {
	Decimals uint64 `bson:"decimals"`
	Value    string `bson:"value"`
}

type TokenPriceDocument struct {
	Usd float64 `bson:"usd"`
}

type TokenDocument struct {
	Decimals        uint64                  `bson:"decimals"`
	Name            string                  `bson:"name"`
	Symbol          string                  `bson:"symbol"`
	GasPriceInUnits GasPriceInUnitsDocument `bson:"gas_price_in_units"`
	GasPriceGwei    string                  `bson:"gas_price_gwei"`
	GasPrice        string                  `bson:"gas_price"`
	ContractAddress string                  `bson:"contract_address"`
	TokenPrice      TokenPriceDocument      `bson:"token_price"`
}

type ExpressGasPriceRateDocument struct {
	AxelarToken            TokenDocument `bson:"axelar_token"`
	DestinationNativeToken TokenDocument `bson:"destination_native_token"`
	EthereumToken          TokenDocument `bson:"ethereum_token"`
	SourceToken            TokenDocument `bson:"source_token"`
}

type ExpressFeeDocument struct {
	ExpressGasOverheadFee    float64 `bson:"express_gas_overhead_fee"`
	ExpressGasOverheadFeeUsd float64 `bson:"express_gas_overhead_fee_usd"`
	RelayerFee               float64 `bson:"relayer_fee"`
	RelayerFeeUsd            float64 `bson:"relayer_fee_usd"`
	Total                    float64 `bson:"total"`
	TotalUsd                 float64 `bson:"total_usd"`
}

type FeesDocument struct {
	AxelarToken                 TokenDocument      `bson:"axelar_token"`
	BaseFee                     float64            `bson:"base_fee"`
	DestinationBaseFee          float64            `bson:"destination_base_fee"`
	DestinationBaseFeeString    string             `bson:"destination_base_fee_string"`
	DestinationBaseFeeUsd       float64            `bson:"destination_base_fee_usd"`
	DestinationConfirmFee       float64            `bson:"destination_confirm_fee"`
	DestinationExpressFee       ExpressFeeDocument `bson:"destination_express_fee"`
	DestinationNativeToken      TokenDocument      `bson:"destination_native_token"`
	EthereumToken               TokenDocument      `bson:"ethereum_token"`
	ExecuteGasMultiplier        float64            `bson:"execute_gas_multiplier"`
	ExecuteMinGasPrice          string             `bson:"execute_min_gas_price"`
	ExpressExecuteGasMultiplier float64            `bson:"express_execute_gas_multiplier"`
	ExpressFee                  float64            `bson:"express_fee"`
	ExpressFeeString            string             `bson:"express_fee_string"`
	ExpressFeeUsd               float64            `bson:"express_fee_usd"`
	ExpressSupported            bool               `bson:"express_supported"`
	SourceBaseFee               float64            `bson:"source_base_fee"`
	SourceBaseFeeUsd            float64            `bson:"source_base_fee_usd"`
	SourceConfirmFee            float64            `bson:"source_confirm_fee"`
	SourceExpressFee            ExpressFeeDocument `bson:"source_express_fee"`
	SourceToken                 TokenDocument      `bson:"source_token"`
}

type GasDocument struct {
	GasExecuteAmount    float64 `bson:"gas_execute_amount"`
	GasApproveAmount    float64 `bson:"gas_approve_amount"`
	GasExpressFeeAmount float64 `bson:"gas_express_fee_amount"`
	GasUsedAmount       float64 `bson:"gas_used_amount"`
	GasRemainAmount     float64 `bson:"gas_remain_amount"`
	GasPaidAmount       float64 `bson:"gas_paid_amount"`
	GasBaseFeeAmount    float64 `bson:"gas_base_fee_amount"`
	GasExpressAmount    float64 `bson:"gas_express_amount"`
	GasUsedValue        float64 `bson:"gas_used_value"`
}

type TimeSpentDocument struct {
	CallConfirm      int `bson:"call_confirm"`
	CallApproved     int `bson:"call_approved"`
	Total            int `bson:"total"`
	ApprovedExecuted int `bson:"approved_executed"`
}

type GMPDocument struct {
	Approved                              GMPStepDocument             `bson:"approved"`
	Call                                  GMPStepDocument             `bson:"call"`
	Confirm                               ConfirmDocument             `bson:"confirm"`
	Executed                              GMPStepDocument             `bson:"executed"`
	ExpressGasPriceRate                   ExpressGasPriceRateDocument `bson:"express_gas_price_rate"`
	Fees                                  FeesDocument                `bson:"fees"`
	Gas                                   GasDocument                 `bson:"gas"`
	GasPaid                               GMPStepDocument             `bson:"gas_paid"`
	GasPriceRate                          ExpressGasPriceRateDocument `bson:"gas_price_rate"`
	Refunded                              GMPStepDocument             `bson:"refunded"`
	IsInvalidPayloadHash                  bool                        `bson:"is_invalid_payload_hash"`
	CommandID                             string                      `bson:"command_id"`
	IsInvalidSourceAddress                bool                        `bson:"is_invalid_source_address"`
	IsInvalidContractAddress              bool                        `bson:"is_invalid_contract_address"`
	IsInvalidDestinationChain             bool                        `bson:"is_invalid_destination_chain"`
	IsCallFromRelayer                     bool                        `bson:"is_call_from_relayer"`
	IsInvalidSymbol                       bool                        `bson:"is_invalid_symbol"`
	IsInvalidAmount                       bool                        `bson:"is_invalid_amount"`
	IsInvalidCall                         bool                        `bson:"is_invalid_call"`
	TimeSpent                             TimeSpentDocument           `bson:"time_spent"`
	IsInvalidGasPaid                      bool                        `bson:"is_invalid_gas_paid"`
	IsInvalidGasPaidMismatchSourceAddress bool                        `bson:"is_invalid_gas_paid_mismatch_source_address"`
	IsInsufficientFee                     bool                        `bson:"is_insufficient_fee"`
	ConfirmFailed                         bool                        `bson:"confirm_failed"`
	ConfirmFailedEvent                    *string                     `bson:"confirm_failed_event,omitempty"`
	ExecutingAt                           int64                       `bson:"executing_at"`
	IsNotEnoughGas                        bool                        `bson:"is_not_enough_gas"`
	ExecutePendingTransactionHash         *string                     `bson:"execute_pending_transaction_hash,omitempty"`
	NotEnoughGasToExecute                 bool                        `bson:"not_enough_gas_to_execute"`
	ExecuteNonce                          *string                     `bson:"execute_nonce,omitempty"`
	RefundingAt                           int64                       `bson:"refunding_at"`
	ToRefund                              bool                        `bson:"to_refund"`
	IsExecuteFromRelayer                  bool                        `bson:"is_execute_from_relayer"`
	RefundNonce                           *string                     `bson:"refund_nonce,omitempty"`
	ID                                    string                      `bson:"id"`
	Status                                string                      `bson:"status"`
	SimplifiedStatus                      string                      `bson:"simplified_status"`
	GasStatus                             string                      `bson:"gas_status"`
	IsTwoWay                              bool                        `bson:"is_two_way"`
}
