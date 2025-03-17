package creator_service

import (
	"bridge-accounts-creator/src/config"
	"bridge-accounts-creator/src/models/account_model"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreatorService interface {
	CreateAccountService() (*account_model.Account, error)
}

type creatorService struct {
	conf *config.Config
}

func NewCreatorService(conf *config.Config) CreatorService {
	return &creatorService{
		conf: conf,
	}
}

func (s *creatorService) CreateAccountService() (*account_model.Account, error) {
	fmt.Printf("\nfmt::Creator::CreateAccountService\n")

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to generate private key: %v\n", err))
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("could not cast public key to ECDSA")
	}
	publicKeyStr := hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	account := account_model.Account{
		PrivateKey: privateKeyStr,
		PublicKey:  publicKeyStr,
		Address:    address,
	}

	fmt.Printf("\nfmt::Creator::CreateAccountService::account: %v\n", account)

	return &account, nil
}
