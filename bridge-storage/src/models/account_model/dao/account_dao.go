package dao

import (
	"bridge-storage/src/models/account_model"
	"bridge-storage/src/models/account_model/dto"
	"fmt"
	"gorm.io/gorm"
)

type AccountDAO interface {
	Create(accountToCreate dto.AccountToCreate) (*account_model.Account, error)
}

type accountDAO struct {
	DB *gorm.DB
}

func NewAccountDAO(db *gorm.DB) AccountDAO {
	return &accountDAO{
		DB: db,
	}
}

func (d *accountDAO) Create(accountToCreate dto.AccountToCreate) (*account_model.Account, error) {
	fmt.Printf("Storage::AccountDAO::Create::accountToCreate: %+v\n", accountToCreate)
	account := account_model.Account{
		PublicKey:  accountToCreate.PublicKey,
		PrivateKey: accountToCreate.PrivateKey,
		Address:    accountToCreate.Address,
	}

	fmt.Printf("Storage::AccountDAO::Create::account: %+v\n", account)

	result := d.DB.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}
