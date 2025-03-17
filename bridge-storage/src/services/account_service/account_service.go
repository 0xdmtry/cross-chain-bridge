package account_service

import (
	"bridge-storage/src/config"
	"bridge-storage/src/models/account_model"
	account_dao "bridge-storage/src/models/account_model/dao"
	account_dto "bridge-storage/src/models/account_model/dto"
	"fmt"
)

type AccountService interface {
	CreateAccountService(publicKey string, privateKey string, address string) (*account_model.Account, error)
}

type accountService struct {
	accountDAO account_dao.AccountDAO
	conf       *config.Config
}

func NewAccountService(accountDAO account_dao.AccountDAO, conf *config.Config) AccountService {
	return &accountService{
		accountDAO: accountDAO,
		conf:       conf,
	}
}

func (s *accountService) CreateAccountService(publicKey string, privateKey string, address string) (*account_model.Account, error) {
	fmt.Printf("Storage::AccountService::CreateAccountService::publicKey: %+v\n", publicKey)
	fmt.Printf("Storage::AccountService::CreateAccountService::privateKey: %+v\n", privateKey)
	fmt.Printf("Storage::AccountService::CreateAccountService::address: %+v\n", address)
	accountDto := account_dto.AccountToCreate{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Address:    address,
	}

	account, err := s.accountDAO.Create(accountDto)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Storage::AccountService::CreateAccountService::account: %+v\n", account)

	return account, nil
}
