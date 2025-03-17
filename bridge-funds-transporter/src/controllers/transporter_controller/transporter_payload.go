package transporter_controller

import "math/big"

type TransactionPayload struct {
	ChainUrl         string   `json:"chainUrl"`
	PrivateKeyStr    string   `json:"privateKeyStr"`
	RecipientAddress string   `json:"recipientAddress"`
	Amount           *big.Int `json:"amount"`
	ChainID          *big.Int `json:"chainID"`
	GasLimit         uint64   `json:"gasLimit"`
}
