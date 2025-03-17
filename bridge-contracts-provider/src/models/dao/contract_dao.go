package dao

import (
	"bridge-contracts-provider/src/config"
	"bridge-contracts-provider/src/helpers/logger"
	"bridge-contracts-provider/src/models/dto"
	"encoding/json"
	"io"
	"os"
)

type ContractDAO interface {
	GetContracts() ([]dto.ContractDTO, error)
}

type contractDAO struct {
	conf *config.Config
}

func NewContractDAO(conf *config.Config) ContractDAO {
	return &contractDAO{
		conf: conf,
	}
}

func (c *contractDAO) GetContracts() ([]dto.ContractDTO, error) {

	file, err := os.Open(c.conf.ContractsSource)
	if err != nil {
		logger.Error("ContractsProvider::GetContracts::os.Open:", err)
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		logger.Error("ContractsProvider::GetContracts::io.ReadAll:", err)
		return nil, err
	}

	var contractDTO []dto.ContractDTO
	if err := json.Unmarshal(byteValue, &contractDTO); err != nil {
		logger.Error("ContractsProvider::GetContracts::json.Unmarshal:", err)
		return nil, err
	}

	return contractDTO, nil
}
