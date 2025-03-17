package dto

type ContractDTO struct {
	Endpoint     string `json:"endpoint"`
	Target       string `json:"target"`
	Network      string `json:"network"`
	ChainId      int64  `json:"chainId"`
	WalletKey    string `json:"walletKey"`
	IsEth        bool   `json:"isEth"`
	IsCompilable bool   `json:"isCompilable"`
	Name         string `json:"name"`
	CompilerPath string `json:"compilerPath"`
	DeployerPath string `json:"deployerPath"`
	ContractPath string `json:"contractPath"`
	OutputPath   string `json:"outputPath"`
}

type ContractToCompileDTO struct {
	Endpoint             string `json:"endpoint"`
	Target               string `json:"target"`
	Network              string `json:"network"`
	ChainId              int64  `json:"chainId"`
	WalletKey            string `json:"walletKey"`
	IsEth                bool   `json:"isEth"`
	IsCompilable         bool   `json:"isCompilable"`
	IsCompiled           bool   `json:"isCompiled"`
	Name                 string `json:"name"`
	CompilerPath         string `json:"compilerPath"`
	DeployerPath         string `json:"deployerPath"`
	ContractPath         string `json:"contractPath"`
	OutputPath           string `json:"outputPath"`
	CompiledContractPath string `json:"compiledContractPath"`
}
