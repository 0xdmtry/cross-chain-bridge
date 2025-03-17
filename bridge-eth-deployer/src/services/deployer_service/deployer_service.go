package deployer_service

import (
	"bridge-eth-deployer/src/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strings"
)

type DeployerService interface {
	DeployContract(path, name, endpoint, walletKey string, chainId int64) error
	getAbiPath(path, name string) string
	getBinPath(path, name string) string
}

type deployerService struct {
	conf *config.Config
}

func NewDeployerService(conf *config.Config) DeployerService {
	return &deployerService{
		conf: conf,
	}
}

func (s *deployerService) DeployContract(path, name, endpoint, walletKey string, chainId int64) error {

	// Connect to Ethereum client
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		fmt.Printf("ERROR::EthDeployer::DeployContract::ethclient.Dial: %v\n", err)
		return err
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(walletKey)
	if err != nil {
		fmt.Printf("ERROR::Fmt::EthDeployer::DeployContract::crypto.HexToECDSA: %v\n", err)
		return err
	}

	// Create auth
	chainID := big.NewInt(chainId)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Printf("ERROR::EthDeployer::DeployContract::bind.NewKeyedTransactorWithChainID: %v\n", err)
		return err
	}

	// Parse ABI
	abiPath := s.getAbiPath(path, name)

	if !s.fileExists(abiPath) {
		fmt.Printf("ERROR::EthDeployer::DeployContract::abiPath: File does not exist")
		return err
	}

	// Read the ABI file
	abiFileContent, err := os.ReadFile(abiPath)
	if err != nil {
		fmt.Printf("ERROR::abiFileContent %v\n", err)
		return err
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiFileContent)))

	if err != nil {
		fmt.Printf("EthDeployer::DeployContract::abi.JSON:%v\n", err)
		return err
	}

	// Contract bytecode
	binPath := s.getBinPath(path, name)

	if !s.fileExists(binPath) {
		fmt.Printf("ERROR::EthDeployer::DeployContract::binPath: File does not exist")
		return err
	}

	binFileContent, err := os.ReadFile(binPath)
	if err != nil {
		fmt.Printf("ERROR::binFileContent %v\n", err)
		return err
	}

	bytecode := common.FromHex(string(binFileContent))

	//Deploy contract
	_, _, _, err = bind.DeployContract(auth, parsedABI, bytecode, client)
	if err != nil {
		fmt.Printf("ERROR::EthDeployer::DeployContract::bind.DeployContract: %v\n", err)
		return err
	}

	return nil
}

func (s *deployerService) getAbiPath(path, name string) string {
	path = s.trimStr(path)
	path = fmt.Sprintf("%s/%s.abi", path, name)
	path = strings.ReplaceAll(path, "##", "..")
	return path
}

func (s *deployerService) getBinPath(path, name string) string {
	path = s.trimStr(path)
	path = fmt.Sprintf("%s/%s.bin", path, name)
	path = strings.ReplaceAll(path, "##", "..")
	return path
}

func (s *deployerService) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (s *deployerService) trimStr(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "..", "##")

	return str
}
