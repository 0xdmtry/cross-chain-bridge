package transporter_service

import (
	"bridge-funds-transporter/src/config"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type TransporterService interface {
	TransferFunds(chainUrl, privateKeyStr, recipientAddress string, amount *big.Int, chainID *big.Int, gasLimit uint64) error
}

type transporterService struct {
	conf *config.Config
}

func NewTransporterService(conf *config.Config) TransporterService {
	return &transporterService{
		conf: conf,
	}
}

func (s *transporterService) TransferFunds(chainUrl, privateKeyStr, recipientAddress string, amount *big.Int, chainID *big.Int, gasLimit uint64) error {
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Printf("Failed to create private key: %v", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Error casting public key to ECDSA\n")
		return errors.New("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Printf("FundsTransporter::TransferFunds::nonce::err: %v", err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("FundsTransporter::TransferFunds::gasPrice::err: %v", err)
		return err
	}

	toAddress := common.HexToAddress(recipientAddress)
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Printf("FundsTransporter::TransferFunds::signedTx::err: %v", err)
		return err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}

	fmt.Printf("Tx sent: %s", signedTx.Hash().Hex())

	return nil
}
