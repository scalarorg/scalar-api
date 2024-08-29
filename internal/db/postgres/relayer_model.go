package postgres

import (
	"database/sql"
)

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
	ContractAddress      sql.NullString `gorm:"column:contractAddress"`
	Payload              sql.NullString
	PayloadHash          sql.NullString
	SourceAddress        sql.NullString
	DestUserAddress      sql.NullString
	DestContractAddress  sql.NullString //Same as ContractAddress
	Amount               sql.NullString
	ContractCallApproved ContractCallApproved `gorm:"foreignKey:ID;references:ID"`
	RelayDataID          uint                 `gorm:"foreignKey:ID;references:ID"`
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
