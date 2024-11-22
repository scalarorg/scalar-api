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

func ToReadableStatus(status int) RelayDataStatus {
	switch status {
	case 0:
		return Pending
	case 1:
		return Approved
	case 2:
		return Success
	case 3:
		return Failed
	default:
		return Undefined
	}
}

type RelayData struct {
	ID                      string `gorm:"primaryKey"`
	PacketSequence          sql.NullInt32
	ExecuteHash             sql.NullString
	Status                  sql.NullInt32
	From                    sql.NullString
	To                      sql.NullString
	CreatedAt               sql.NullTime
	UpdatedAt               sql.NullTime
	ContractCallID          sql.NullInt32 `gorm:"foreignKey:ID;references:ID"`
	ContractCall            ContractCall  `gorm:"foreignKey:ID;references:ID"`
	ContractCallWithTokenID sql.NullInt32 `gorm:"foreignKey:ID;references:ID"`
	//ContractCallWithToken   ContractCallWithToken `gorm:"foreignKey:ID;references:ID"`
}
type ContractCall struct {
	ID                   string `gorm:"primaryKey"`
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
	BlockNumber          sql.NullInt32 `gorm:"column:block_number"`
	LogIndex             sql.NullInt32
	ContractAddress      sql.NullString
	Payload              sql.NullString
	PayloadHash          sql.NullString
	SourceAddress        sql.NullString
	DestUserAddress      sql.NullString
	DestContractAddress  sql.NullString //Same as ContractAddress
	Amount               sql.NullString
	ContractCallApproved ContractCallApproved `gorm:"foreignKey:ID;references:ID"`
	CommandExecuted      CommandExecuted      `gorm:"foreignKey:ID;references:ID"`
	StakerPublicKey      sql.NullString
	SenderAddress        sql.NullString
	RelayDataID          string
	TxHash               sql.NullString
	TxHex                []byte
}

type ContractCallWithToken struct {
	ID                  string `gorm:"primaryKey"`
	CreatedAt           sql.NullTime
	UpdatedAt           sql.NullTime
	BlockNumber         sql.NullInt32
	ContractAddress     sql.NullString
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
	SourceChain      sql.NullString
	DestinationChain sql.NullString
	TxHash           sql.NullString
	BlockNumber      sql.NullInt32
	LogIndex         sql.NullInt32
	SourceAddress    sql.NullString
	ContractAddress  sql.NullString
	SourceTxHash     sql.NullString
	SourceEventIndex sql.NullInt64
	PayloadHash      sql.NullString
	CommandId        sql.NullString
}

type CommandExecuted struct {
	ID               string `gorm:"primaryKey"`
	SourceChain      sql.NullString
	DestinationChain sql.NullString
	TxHash           sql.NullString
	BlockNumber      sql.NullInt32
	LogIndex         sql.NullInt32
	CommandId        sql.NullString
	Status           sql.NullInt32
	ReferenceTxHash  sql.NullString
	Amount           sql.NullString
}

type ContractCallWithTokenApproved struct {
	ID               string `gorm:"primaryKey"`
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	SourceChain      sql.NullString
	DestinationChain sql.NullString
	TxHash           sql.NullString
	BlockNumber      sql.NullInt32
	LogIndex         sql.NullInt32
	SourceAddress    sql.NullString
	ContractAddress  sql.NullString
	SourceTxHash     sql.NullString
	SourceEventIndex sql.NullInt64
	PayloadHash      sql.NullString
	CommandId        sql.NullString
}

func (RelayData) TableName() string {
	return "relay_data"
}
