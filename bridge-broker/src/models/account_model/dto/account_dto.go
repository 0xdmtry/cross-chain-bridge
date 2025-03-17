package dto

import "bridge-broker/src/models/account_model"

type CreatedAccountResponse struct {
	Account account_model.CreatedAccount `json:"account"`
}

type PreservedAccountResponse struct {
	Account account_model.PreservedAccount `json:"account"`
}

type TransactionResponse struct {
	Transaction account_model.Transaction `json:"transaction"`
}
