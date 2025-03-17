package broker_service

import (
	"bridge-broker/src/config"
	"bridge-broker/src/helpers/logger"
	"bridge-broker/src/models/account_model"
	account_dto "bridge-broker/src/models/account_model/dto"
	contract_dto "bridge-broker/src/models/contract_model/dto"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type BrokerService interface {
	CreateAccountService() (*account_model.PreservedAccount, error)

	preserveAccountService(account account_model.CreatedAccount) (*account_model.PreservedAccount, error)
	createAccountService() (*account_model.CreatedAccount, error)

	ProcessContractService() error
	getContractsInfo() ([]contract_dto.ContractToCompileDTO, error)
	compileContracts(contractsInfo []contract_dto.ContractToCompileDTO)
	getCompilerPath(compilerPath, contractPath, outputPath, name string) string
}

type brokerService struct {
	conf *config.Config
}

func NewBrokerService(conf *config.Config) BrokerService {
	return &brokerService{
		conf: conf,
	}
}

func (s *brokerService) CreateAccountService() (*account_model.PreservedAccount, error) {
	fmt.Printf("\nfmt::Broker::CreateAccountService::s.conf.AccountCreatorUrl: %v\n", s.conf.AccountCreatorUrl)

	createdAccount, err := s.createAccountService()
	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountService::ERROR: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::CreateAccountService::createdAccount: %v\n", createdAccount)

	preservedAccount, err := s.preserveAccountService(*createdAccount)
	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountService::ERROR: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::CreateAccountService::preservedAccount: %v\n", preservedAccount)

	return preservedAccount, nil
}

func (s *brokerService) transferFundsService(preservedAccount *account_model.PreservedAccount) (*account_model.Transaction, error) {

	accountData := s.getTransactionData(s.conf.ChainUrl, s.conf.privateKeyStr, preservedAccount.Address, s.conf.AccountCreationAmount, s.conf.ChainID, s.conf.GasLimit)
	resp, err := http.PostForm(s.conf.StorageUrl, accountData)
	if err != nil {
		fmt.Printf("\nfmt::Broker::preserveAccountService::ERROR 1: %v\n", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	fmt.Printf("\nfmt::Broker::preserveAccountService::resp: %v\n", resp)

	return nil, nil
}

func (s *brokerService) preserveAccountService(account account_model.CreatedAccount) (*account_model.PreservedAccount, error) {
	accountData := s.getAccountData(account.PrivateKey, account.PublicKey, account.Address)
	resp, err := http.PostForm(s.conf.StorageUrl, accountData)
	if err != nil {
		fmt.Printf("\nfmt::Broker::preserveAccountService::ERROR 1: %v\n", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	fmt.Printf("\nfmt::Broker::preserveAccountService::resp: %v\n", resp)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nfmt::Broker::preserveAccountService::ERROR 2: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::preserveAccountService::bodyBytes: %v\n", bodyBytes)
	fmt.Printf("\nfmt::Broker::preserveAccountService::string(bodyBytes): %v\n", string(bodyBytes))

	var payload account_dto.PreservedAccountResponse
	err = json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		fmt.Printf("\nfmt::Broker::preserveAccountService::ERROR 3: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::preserveAccountService::CreatedAccount: %v\n", payload.Account)

	return &payload.Account, nil
}

func (s *brokerService) createAccountService() (*account_model.CreatedAccount, error) {
	resp, err := http.Get(fmt.Sprintf("%s", s.conf.AccountCreatorUrl))
	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountService::ERROR 1: %v\n", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	fmt.Printf("\nfmt::Broker::CreateAccountService::resp: %v\n", resp)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountService::ERROR 2: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::CreateAccountService::bodyBytes: %v\n", bodyBytes)
	fmt.Printf("\nfmt::Broker::CreateAccountService::string(bodyBytes): %v\n", string(bodyBytes))

	var payload account_dto.CreatedAccountResponse
	err = json.Unmarshal(bodyBytes, &payload)

	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountService::ERROR 3: %v\n", err)
		return nil, err
	}
	fmt.Printf("\nfmt::Broker::CreateAccountService::CreatedAccount: %v\n", payload.Account)

	return &payload.Account, nil
}

func (s *brokerService) ProcessContractService() error {
	contractsInfo, err := s.getContractsInfo()
	if err != nil {
		logger.Error("Broker::CreateAccounts:", err)
		return err
	}
	s.compileContracts(contractsInfo)

	return nil
}

func (s *brokerService) getContractsInfo() ([]contract_dto.ContractToCompileDTO, error) {
	resp, err := http.Get(fmt.Sprintf("%s", s.conf.ContractsProviderUrl))
	if err != nil {
		logger.Error("Broker::getContractsInfo::http.Get:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Broker::getContractsInfo::io.ReadAll:", err)
		return nil, err
	}

	var response []contract_dto.ContractToCompileDTO
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		logger.Error("Broker::getContractsInfo::Unmarshal:", err)
		return nil, err
	}

	return response, nil
}

func (s *brokerService) compileContracts(contractsInfo []contract_dto.ContractToCompileDTO) {

	for _, contractInfo := range contractsInfo {
		if contractInfo.IsCompilable && contractInfo.IsEth {
			compilerPath := s.getCompilerPath(contractInfo.CompilerPath, contractInfo.ContractPath, contractInfo.OutputPath, contractInfo.Name)
			compilerResp, err := http.Get(fmt.Sprintf("%s", compilerPath))
			if err != nil {
				logger.Error("Broker::compileContracts::http.Get::compilerResp:", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(compilerResp.Body)

			compilerBodyBytes, err := io.ReadAll(compilerResp.Body)
			if err != nil {
				logger.Error("Broker::compileContracts::io.ReadAll:", err)
			}

			deployerData := s.getDeployerData(string(compilerBodyBytes), contractInfo.Name, contractInfo.Endpoint, contractInfo.WalletKey, contractInfo.ChainId)

			deployerResp, err := http.PostForm(contractInfo.DeployerPath, deployerData)

			if err != nil {
				logger.Error("Broker::compileContracts:http.Get::deployerResp:", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(deployerResp.Body)

			_, err = io.ReadAll(deployerResp.Body)
			if err != nil {
				logger.Error("Broker::compileContracts::io.ReadAll", err)
			}

			return
		}
	}
}

func (s *brokerService) getCompilerPath(compilerPath, contractPath, outputPath, name string) string {
	return fmt.Sprintf("%s?path=%s&output=%s&name=%s", compilerPath, contractPath, outputPath, name)
}

func (s *brokerService) getDeployerPath(deployerPath, path, name, endpoint, wallet string, chainId int64) string {
	return fmt.Sprintf("%s?path=%s&name=%s&endpoint=%s&wallet=%s&chain=%v", deployerPath, path, name, endpoint, wallet, chainId)
}

func (s *brokerService) getTransactionData(chainUrl string, privateKeyStr string, recipientAddress string, amount *big.Int, chainID *big.Int, gasLimit uint64) url.Values {
	data := url.Values{}
	data.Set("chainUrl", chainUrl)
	data.Set("privateKeyStr", privateKeyStr)
	data.Set("recipientAddress", recipientAddress)

	if amount != nil {
		data.Set("amount", amount.String())
	} else {
		data.Set("amount", "")
	}

	if chainID != nil {
		data.Set("chainID", chainID.String())
	} else {
		data.Set("chainID", "")
	}

	data.Set("gasLimit", strconv.FormatUint(gasLimit, 10))

	return data
}

func (s *brokerService) getAccountData(privateKey string, publicKey string, address string) url.Values {
	data := url.Values{}
	data.Set("privateKey", privateKey)
	data.Set("publicKey", publicKey)
	data.Set("address", address)
	return data
}

func (s *brokerService) getDeployerData(path string, name string, endpoint string, walletKey string, chainId int64) url.Values {
	data := url.Values{}
	data.Set("path", path)
	data.Set("name", name)
	data.Set("endpoint", endpoint)
	data.Set("walletKey", walletKey)
	data.Set("chainId", strconv.FormatInt(chainId, 10))
	return data
}

func (s *brokerService) trimStr(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, " ", "")

	return str
}
