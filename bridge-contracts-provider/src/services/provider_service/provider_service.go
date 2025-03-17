package provider_service

import (
	"bridge-contracts-provider/src/config"
	"bridge-contracts-provider/src/helpers/logger"
	"bridge-contracts-provider/src/models/dao"
	"bridge-contracts-provider/src/models/dto"
)

type ProviderService interface {
	ProvideContracts() ([]dto.ContractDTO, error)
}

type providerService struct {
	conf        *config.Config
	contractDAO dao.ContractDAO
}

func NewProviderService(conf *config.Config, contractDAO dao.ContractDAO) ProviderService {
	return &providerService{
		conf:        conf,
		contractDAO: contractDAO,
	}
}

func (s *providerService) ProvideContracts() ([]dto.ContractDTO, error) {
	contractsInfo, err := s.contractDAO.GetContracts()
	if err != nil {
		logger.Error("ContractsProvider::ProvideContracts:", err)
		return nil, err
	}
	return contractsInfo, nil
}
