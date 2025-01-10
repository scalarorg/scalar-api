package models

type GMPTx struct {
	ID      uint `gorm:"primaryKey"`
	Hash    string
	BlockId uint
	Code    uint
	Memo    string
}

type TxMessage struct {
	ID            uint `gorm:"primaryKey"`
	TxID          uint `gorm:"uniqueIndex:txMessageIndex,priority:1"`
	MessageID     uint `gorm:"uniqueIndex:txMessageIndex,priority:2"`
	BlockId       uint
	MessageDetail string
	Tx            GMPTx `gorm:"foreignKey:TxID;references:ID"`
}

type CreatedAtDocument struct {
	Week    int64 `json:"week,omitempty"`
	Hour    int64 `json:"hour,omitempty"`
	Month   int64 `json:"month,omitempty"`
	Year    int64 `json:"year,omitempty"`
	Ms      int64 `json:"ms,omitempty"`
	Day     int64 `json:"day,omitempty"`
	Quarter int64 `json:"quarter,omitempty"`
}

type LogDocument struct {
	BlockHash        string   `json:"blockHash,omitempty"`
	Address          string   `json:"address,omitempty"`
	LogIndex         uint     `json:"logIndex,omitempty"`
	Data             string   `json:"data,omitempty"`
	Removed          bool     `json:"removed,omitempty"`
	Topics           []string `json:"topics,omitempty"`
	BlockNumber      uint64   `json:"blockNumber,omitempty"`
	TransactionIndex uint     `json:"transactionIndex,omitempty"`
	TransactionHash  string   `json:"transactionHash,omitempty"`
}

type ReceiptDocument struct {
	BlockHash         string        `json:"blockHash,omitempty"`
	ContractAddress   *string       `json:"contractAddress,omitempty"`
	TransactionIndex  uint          `json:"transactionIndex,omitempty"`
	Type              uint8         `json:"type,omitempty"`
	Confirmations     uint          `json:"confirmations,omitempty"`
	TransactionHash   string        `json:"transactionHash,omitempty"`
	GasUsed           string        `json:"gasUsed,omitempty"`
	BlockNumber       uint64        `json:"blockNumber,omitempty"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed,omitempty"`
	From              string        `json:"from,omitempty"`
	To                string        `json:"to,omitempty"`
	EffectiveGasPrice string        `json:"effectiveGasPrice,omitempty"`
	Logs              []LogDocument `json:"logs,omitempty"`
	Status            uint8         `json:"status,omitempty"`
}

type ReturnValuesDocument struct {
	CommandID                  string `json:"commandId,omitempty"`
	SourceTxHash               string `json:"sourceTxHash,omitempty"`
	SourceAddress              string `json:"sourceAddress,omitempty"`
	SourceChain                string `json:"sourceChain,omitempty"`
	SourceEventIndex           string `json:"sourceEventIndex,omitempty"`
	Sender                     string `json:"sender,omitempty"`
	ContractAddress            string `json:"contractAddress,omitempty"`
	DestinationChain           string `json:"destinationChain,omitempty"`
	DestinationContractAddress string `json:"destinationContractAddress,omitempty"`
	PayloadHash                string `json:"payloadHash,omitempty"`
	Payload                    string `json:"payload,omitempty"`
}

type TransactionDocument struct {
	BlockHash            string `json:"blockHash,omitempty"`
	YParity              string `json:"yParity,omitempty"`
	TransactionIndex     uint   `json:"transactionIndex,omitempty"`
	Type                 uint8  `json:"type,omitempty"`
	Nonce                uint64 `json:"nonce,omitempty"`
	R                    string `json:"r,omitempty"`
	S                    string `json:"s,omitempty"`
	ChainID              int64  `json:"chainId,omitempty"`
	V                    uint64 `json:"v,omitempty"`
	BlockNumber          uint64 `json:"blockNumber,omitempty"`
	Gas                  string `json:"gas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	From                 string `json:"from,omitempty"`
	To                   string `json:"to,omitempty"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	Value                uint64 `json:"value,omitempty"`
	Hash                 string `json:"hash,omitempty"`
	GasPrice             string `json:"gasPrice,omitempty"`
}

type GMPStepDocument struct {
	Chain                string               `json:"chain,omitempty"`
	SourceChain          string               `json:"sourceChain,omitempty"`
	ContractAddress      string               `json:"contract_address,omitempty"`
	Address              string               `json:"address,omitempty"`
	Topics               []string             `json:"topics,omitempty"`
	BlockNumber          uint64               `json:"blockNumber"`
	TransactionHash      string               `json:"transactionHash,omitempty"`
	TransactionIndex     uint                 `json:"transactionIndex,omitempty"`
	BlockHash            string               `json:"blockHash,omitempty"`
	LogIndex             uint                 `json:"logIndex"`
	Removed              bool                 `json:"removed,omitempty"`
	ID                   string               `json:"id,omitempty"`
	Event                string               `json:"event,omitempty"`
	EventSignature       string               `json:"eventSignature,omitempty"`
	ReturnValues         ReturnValuesDocument `json:"returnValues,omitempty"`
	ChainType            string               `json:"chain_type,omitempty"`
	DestinationChainType string               `json:"destination_chain_type,omitempty"`
	CreatedAt            CreatedAtDocument    `json:"created_at,omitempty"`
	EventIndex           uint                 `json:"eventIndex,omitempty"`
	BlockTimestamp       int64                `json:"block_timestamp"`
	Receipt              ReceiptDocument      `json:"receipt,omitempty"`
	Transaction          TransactionDocument  `json:"transaction,omitempty"`
	LogIndexAlias        uint                 `json:"_logIndex"`
	TypeAlias            string               `json:"_type,omitempty"`
	ProposalId           string               `json:"proposal_id,omitempty"`
}

type ConfirmDocument struct {
	BlockNumber           uint64 `json:"blockNumber,omitempty"`
	BlockTimestamp        int64  `json:"block_timestamp,omitempty"`
	ConfirmaionTxHash     string `json:"confirmation_txhash,omitempty"`
	Event                 string `json:"event,omitempty"`
	PollId                string `json:"poll_id,omitempty"`
	SourceChain           string `json:"sourceChain,omitempty"`
	SourceTransactionHash string `json:"sourceTransactionHash,omitempty"`
	TransactionHash       string `json:"transactionHash,omitempty"`
	TransactionIndex      uint64 `json:"transactionIndex,omitempty"`
}

type GasPriceInUnitsDocument struct {
	Decimals uint64 `json:"decimals,omitempty"`
	Value    string `json:"value,omitempty"`
}

type TokenPriceDocument struct {
	Usd float64 `json:"usd,omitempty"`
}
type TimeSpentDocument struct {
	CallConfirm      int `json:"call_confirm,omitempty"`
	CallApproved     int `json:"call_approved,omitempty"`
	Total            int `json:"total,omitempty"`
	ApprovedExecuted int `json:"approved_executed,omitempty"`
}

type TokenDocument struct {
	Decimals        uint64                  `json:"decimals,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Symbol          string                  `json:"symbol,omitempty"`
	GasPriceInUnits GasPriceInUnitsDocument `json:"gas_price_in_units,omitempty"`
	GasPriceGwei    string                  `json:"gas_price_gwei,omitempty"`
	GasPrice        string                  `json:"gas_price,omitempty"`
	ContractAddress string                  `json:"contract_address,omitempty"`
	TokenPrice      TokenPriceDocument      `json:"token_price,omitempty"`
}

type ExpressGasPriceRateDocument struct {
	AxelarToken            TokenDocument `json:"axelar_token,omitempty"`
	DestinationNativeToken TokenDocument `json:"destination_native_token,omitempty"`
	EthereumToken          TokenDocument `json:"ethereum_token,omitempty"`
	SourceToken            TokenDocument `json:"source_token,omitempty"`
}

type ExpressFeeDocument struct {
	ExpressGasOverheadFee    float64 `json:"express_gas_overhead_fee,omitempty"`
	ExpressGasOverheadFeeUsd float64 `json:"express_gas_overhead_fee_usd,omitempty"`
	RelayerFee               float64 `json:"relayer_fee,omitempty"`
	RelayerFeeUsd            float64 `json:"relayer_fee_usd,omitempty"`
	Total                    float64 `json:"total,omitempty"`
	TotalUsd                 float64 `json:"total_usd,omitempty"`
}

type FeesDocument struct {
	AxelarToken                 TokenDocument      `json:"axelar_token,omitempty"`
	BaseFee                     float64            `json:"base_fee,omitempty"`
	DestinationBaseFee          float64            `json:"destination_base_fee,omitempty"`
	DestinationBaseFeeString    string             `json:"destination_base_fee_string,omitempty"`
	DestinationBaseFeeUsd       float64            `json:"destination_base_fee_usd,omitempty"`
	DestinationConfirmFee       float64            `json:"destination_confirm_fee,omitempty"`
	DestinationExpressFee       ExpressFeeDocument `json:"destination_express_fee,omitempty"`
	DestinationNativeToken      TokenDocument      `json:"destination_native_token,omitempty"`
	EthereumToken               TokenDocument      `json:"ethereum_token,omitempty"`
	ExecuteGasMultiplier        float64            `json:"execute_gas_multiplier,omitempty"`
	ExecuteMinGasPrice          string             `json:"execute_min_gas_price,omitempty"`
	ExpressExecuteGasMultiplier float64            `json:"express_execute_gas_multiplier,omitempty"`
	ExpressFee                  float64            `json:"express_fee,omitempty"`
	ExpressFeeString            string             `json:"express_fee_string,omitempty"`
	ExpressFeeUsd               float64            `json:"express_fee_usd,omitempty"`
	ExpressSupported            bool               `json:"express_supported,omitempty"`
	SourceBaseFee               float64            `json:"source_base_fee,omitempty"`
	SourceBaseFeeUsd            float64            `json:"source_base_fee_usd,omitempty"`
	SourceConfirmFee            float64            `json:"source_confirm_fee,omitempty"`
	SourceExpressFee            ExpressFeeDocument `json:"source_express_fee,omitempty"`
	SourceToken                 TokenDocument      `json:"source_token,omitempty"`
}

type GasDocument struct {
	GasExecuteAmount    float64 `json:"gas_execute_amount,omitempty"`
	GasApproveAmount    float64 `json:"gas_approve_amount,omitempty"`
	GasExpressFeeAmount float64 `json:"gas_express_fee_amount,omitempty"`
	GasUsedAmount       float64 `json:"gas_used_amount,omitempty"`
	GasRemainAmount     float64 `json:"gas_remain_amount,omitempty"`
	GasPaidAmount       float64 `json:"gas_paid_amount,omitempty"`
	GasBaseFeeAmount    float64 `json:"gas_base_fee_amount,omitempty"`
	GasExpressAmount    float64 `json:"gas_express_amount,omitempty"`
	GasUsedValue        float64 `json:"gas_used_value,omitempty"`
}

type GMPDocument struct {
	Approved                              GMPStepDocument             `json:"approved,omitempty"`
	Call                                  GMPStepDocument             `json:"call,omitempty"`
	Confirm                               ConfirmDocument             `json:"confirm,omitempty"`
	Executed                              GMPStepDocument             `json:"executed,omitempty"`
	ExpressGasPriceRate                   ExpressGasPriceRateDocument `json:"express_gas_price_rate,omitempty"`
	Fees                                  FeesDocument                `json:"fees,omitempty"`
	Gas                                   GasDocument                 `json:"gas,omitempty"`
	GasPaid                               GMPStepDocument             `json:"gas_paid,omitempty"`
	GasPriceRate                          ExpressGasPriceRateDocument `json:"gas_price_rate,omitempty"`
	Refunded                              GMPStepDocument             `json:"refunded,omitempty"`
	IsInvalidPayloadHash                  bool                        `json:"is_invalid_payload_hash,omitempty"`
	CommandID                             string                      `json:"command_id,omitempty"`
	IsInvalidSourceAddress                bool                        `json:"is_invalid_source_address,omitempty"`
	IsInvalidContractAddress              bool                        `json:"is_invalid_contract_address,omitempty"`
	IsInvalidDestinationChain             bool                        `json:"is_invalid_destination_chain,omitempty"`
	IsCallFromRelayer                     bool                        `json:"is_call_from_relayer,omitempty"`
	IsInvalidSymbol                       bool                        `json:"is_invalid_symbol,omitempty"`
	IsInvalidAmount                       bool                        `json:"is_invalid_amount,omitempty"`
	IsInvalidCall                         bool                        `json:"is_invalid_call,omitempty"`
	TimeSpent                             TimeSpentDocument           `json:"time_spent,omitempty"`
	IsInvalidGasPaid                      bool                        `json:"is_invalid_gas_paid,omitempty"`
	IsInvalidGasPaidMismatchSourceAddress bool                        `json:"is_invalid_gas_paid_mismatch_source_address,omitempty"`
	IsInsufficientFee                     bool                        `json:"is_insufficient_fee,omitempty"`
	ConfirmFailed                         bool                        `json:"confirm_failed,omitempty"`
	ConfirmFailedEvent                    *string                     `json:"confirm_failed_event,omitempty,omitempty"`
	ExecutingAt                           int64                       `json:"executing_at,omitempty"`
	IsNotEnoughGas                        bool                        `json:"is_not_enough_gas,omitempty"`
	ExecutePendingTransactionHash         *string                     `json:"execute_pending_transaction_hash,omitempty,omitempty"`
	NotEnoughGasToExecute                 bool                        `json:"not_enough_gas_to_execute,omitempty"`
	ExecuteNonce                          *string                     `json:"execute_nonce,omitempty,omitempty"`
	RefundingAt                           int64                       `json:"refunding_at,omitempty"`
	ToRefund                              bool                        `json:"to_refund,omitempty"`
	IsExecuteFromRelayer                  bool                        `json:"is_execute_from_relayer,omitempty"`
	RefundNonce                           *string                     `json:"refund_nonce,omitempty,omitempty"`
	ID                                    string                      `json:"id,omitempty"`
	Status                                string                      `json:"status,omitempty"`
	SimplifiedStatus                      string                      `json:"simplified_status,omitempty"`
	GasStatus                             string                      `json:"gas_status,omitempty"`
	IsTwoWay                              bool                        `json:"is_two_way,omitempty"`
	CreatedAtDocument
}

type MapBlockEventAttributes map[string]string

type Options struct {
	Size         int
	Offset       int
	EventId      string
	EventType    string
	EventTypes   []string
	StakerPubkey string
}
