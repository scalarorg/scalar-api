package models

type Custodial struct {
	ID              uint   `gorm:"primaryKey;column:id"`
	Name            string `gorm:"column:name"`
	BtcPublicKeyHex string `gorm:"uniqueIndex; not null;column:btc_public_key_hex"`
}

type CustodialGroup struct {
	ID         uint        `gorm:"primaryKey;column:id"`
	Name       string      `gorm:"column:name"`
	BtcAddress string      `gorm:"uniqueIndex; not null;column:taproot_address"` // Calculate from BtcPublicKeyHex of each Custodials
	Quorum     uint        `gorm:"not null;column:quorum"`
	Custodials []Custodial `gorm:"many2many:custodial_group_members;"`
}

type ShortenCustodialGroup struct {
	ID         uint
	Name       string
	BtcAddress string
}
