package account_model

type Account struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
}
