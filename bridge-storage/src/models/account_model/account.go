package account_model

import (
	"bridge-storage/src/models/defaults"
)

type Account struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	PublicKey  string `json:"publicKey" gorm:"unique;not null"`
	PrivateKey string `json:"privateKey" gorm:"unique;not null"`
	Address    string `json:"address" gorm:"unique;not null"`
	defaults.Timestamps
}
