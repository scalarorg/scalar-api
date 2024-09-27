package postgres

import (
	"database/sql"
)

type RelayDataStatus string

const (
	Pending   RelayDataStatus = "Pending"
	Approved  RelayDataStatus = "Approved"
	Success   RelayDataStatus = "Success"
	Failed    RelayDataStatus = "Failed"
	Undefined RelayDataStatus = "Undefined"
)

func ToReadableStatus(status string) RelayDataStatus {
	switch status {
	case "0":
		return Pending
	case "1":
		return Approved
	case "2":
		return Success
	case "3":
		return Failed
	default:
		return Undefined
	}
}

type RelayData struct {
	ID                      string `gorm:"primaryKey"`
	PacketSequence          sql.NullInt32
	ExecuteHash             sql.NullString
	Status                  sql.NullString
	From                    sql.NullString
	To                      sql.NullString
	CreatedAt               sql.NullTime
	UpdatedAt               sql.NullTime
	ContractCallID          sql.NullInt32         `gorm:"foreignKey:ID;references:ID"`
	ContractCall            ContractCall          `gorm:"foreignKey:ID;references:ID"`
	ContractCallWithTokenID sql.NullInt32         `gorm:"foreignKey:ID;references:ID"`
	ContractCallWithToken   ContractCallWithToken `gorm:"foreignKey:ID;references:ID"`
}
type ContractCall struct {
	ID                   string `gorm:"primaryKey"`
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
	BlockNumber          sql.NullInt32  `gorm:"column:blockNumber"`
	LogIndex             sql.NullInt32  `gorm:"column:logIndex"`
	ContractAddress      sql.NullString `gorm:"column:contractAddress"`
	Payload              sql.NullString
	PayloadHash          sql.NullString
	SourceAddress        sql.NullString
	DestUserAddress      sql.NullString
	DestContractAddress  sql.NullString //Same as ContractAddress
	Amount               sql.NullString
	ContractCallApproved ContractCallApproved `gorm:"foreignKey:ID;references:ID"`
	CommandExecuted      CommandExecuted
	StakerPublicKey      sql.NullString `gorm:"column:stakerPublicKey"`
	SenderAddress        sql.NullString `gorm:"column:senderAddress"`
	RelayDataID          uint           `gorm:"foreignKey:ID;references:ID"`
	TxHash               sql.NullString `gorm:"column:txHash"`
	TxHex                []byte         `gorm:"column:txHex"`
}

type ContractCallWithToken struct {
	ID                  string `gorm:"primaryKey"`
	CreatedAt           sql.NullTime
	UpdatedAt           sql.NullTime
	BlockNumber         sql.NullInt32  `gorm:"column:blockNumber"`
	ContractAddress     sql.NullString `gorm:"column:contractAddress"`
	Payload             sql.NullString
	PayloadHash         sql.NullString
	SourceAddress       sql.NullString
	DestUserAddress     sql.NullString
	DestContractAddress sql.NullString
	Amount              sql.NullString
	Symbol              sql.NullString
	RelayDataID         uint `gorm:"foreignKey:ID;references:ID"`
}

type ContractCallApproved struct {
	ID               string `gorm:"primaryKey"`
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	SourceChain      sql.NullString `gorm:"column:sourceChain"`
	DestinationChain sql.NullString `gorm:"column:destinationChain"`
	TxHash           sql.NullString `gorm:"column:txHash"`
	BlockNumber      sql.NullInt32  `gorm:"column:blockNumber"`
	LogIndex         sql.NullInt32  `gorm:"column:logIndex"`
	SourceAddress    sql.NullString `gorm:"column:sourceAddress"`
	ContractAddress  sql.NullString `gorm:"column:contractAddress"`
	SourceTxHash     sql.NullString `gorm:"column:sourceTxHash"`
	SourceEventIndex sql.NullInt64  `gorm:"column:sourceEventIndex"`
	PayloadHash      sql.NullString `gorm:"column:payloadHash"`
	CommandId        sql.NullString `gorm:"column:commandID"`
}

type CommandExecuted struct {
	ID               string         `gorm:"primaryKey"`
	SourceChain      sql.NullString `gorm:"column:sourceChain"`
	DestinationChain sql.NullString `gorm:"column:destinationChain"`
	TxHash           sql.NullString `gorm:"column:txHash"`
	BlockNumber      sql.NullInt32  `gorm:"column:blockNumber"`
	LogIndex         sql.NullInt32  `gorm:"column:logIndex"`
	CommandId        sql.NullString `gorm:"column:commandID"`
	Status           sql.NullInt32  `gorm:"column:status"`
	ReferenceTxHash  sql.NullString `gorm:"column:referenceTxHash"`
	Amount           sql.NullString `gorm:"column:amount"`
}

type ContractCallWithTokenApproved struct {
	ID               string `gorm:"primaryKey"`
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	SourceChain      sql.NullString `gorm:"column:sourceChain"`
	DestinationChain sql.NullString `gorm:"column:destinationChain"`
	TxHash           sql.NullString `gorm:"column:txHash"`
	BlockNumber      sql.NullInt32  `gorm:"column:blockNumber"`
	LogIndex         sql.NullInt32  `gorm:"column:logIndex"`
	SourceAddress    sql.NullString `gorm:"column:sourceAddress"`
	ContractAddress  sql.NullString `gorm:"column:contractAddress"`
	SourceTxHash     sql.NullString `gorm:"column:sourceTxHash"`
	SourceEventIndex sql.NullInt64  `gorm:"column:sourceEventIndex"`
	PayloadHash      sql.NullString `gorm:"column:payloadHash"`
	CommandId        sql.NullString `gorm:"column:commandID"`
}

func (RelayData) TableName() string {
	return "RelayData"
}

func (ContractCall) TableName() string {
	return "CallContract"
}

func (ContractCallWithToken) TableName() string {
	return "CallContractWithToken"
}

func (ContractCallApproved) TableName() string {
	return "ContractCallApproved"
}

func (ContractCallWithTokenApproved) TableName() string {
	return "ContractCallWithTokenApproved"
}
