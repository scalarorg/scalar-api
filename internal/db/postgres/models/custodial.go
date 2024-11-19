package models

type Custodial struct {
	ID              uint   `gorm:"primaryKey;column:id"`
	Name            string `gorm:"uniqueIndex;column:name"`
	BtcPublicKeyHex string `gorm:"uniqueIndex;column:btc_public_key_hex"`
}

type CustodialGroup struct {
	ID         uint        `gorm:"primaryKey;column:id"`
	Name       string      `gorm:"uniqueIndex;column:name"`
	Quorum     uint        `gorm:"not null;column:quorum"`
	Custodials []Custodial `gorm:"many2many:custodial_group_members;"`
}
