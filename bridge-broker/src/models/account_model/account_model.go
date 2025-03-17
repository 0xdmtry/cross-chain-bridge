package account_model

type CreatedAccount struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
}

type PreservedAccount struct {
	ID         uint   `json:"id"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type Transaction struct {
}
