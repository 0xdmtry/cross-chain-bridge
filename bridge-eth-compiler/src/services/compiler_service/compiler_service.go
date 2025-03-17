package compiler_service

import (
	"bridge-eth-compiler/src/config"
	"bridge-eth-compiler/src/helpers/logger"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type CompilerService interface {
	CompileContract(path, output, name string) (string, error)
	getTimeStamp() string
	getHash() string
	getOutputPath(output, name, timestamp, hash string) string
	createDirectory(path string) error
	execCommand(outputPath, path string) error
}

type compilerService struct {
	conf *config.Config
}

func NewCompilerService(conf *config.Config) CompilerService {
	return &compilerService{
		conf: conf,
	}
}

func (s *compilerService) CompileContract(path, output, name string) (string, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Error("EthCompiler::CompileContract::IsNotExist(path):", err)
	}

	if _, err := os.Stat(output); os.IsNotExist(err) {
		logger.Error("EthCompiler::CompileContract::IsNotExist(output):", err)
	}

	timeStamp := s.getTimeStamp()
	hash := s.getHash()
	outputPath := s.getOutputPath(output, name, timeStamp, hash)

	err := s.createDirectory(outputPath)
	if err != nil {
		logger.Error("EthCompiler::CompileContract::createDirectory:", err)
		return "", err
	}

	err = s.execCommand(outputPath, path)
	if err != nil {
		logger.Error("EthCompiler::CompileContract::execCommand:", err)
		return "", err
	}

	return outputPath, nil
}

func (s *compilerService) getTimeStamp() string {
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02-15-04-05.000")

	return formattedTime
}

func (s *compilerService) getHash() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	hasher := sha256.New()
	hasher.Write(randomBytes)

	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	shortHash := hashString[:6]

	return shortHash
}

func (s *compilerService) getOutputPath(output, name, timestamp, hash string) string {
	return fmt.Sprintf("%s%s-%s-%s", output, name, timestamp, hash)
}

func (s *compilerService) createDirectory(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		logger.Error("EthCompiler::CompileContract::createDirectory:", err)
		return err
	}
	return nil
}

func (s *compilerService) execCommand(outputPath, path string) error {
	cmd := exec.Command("solc", "--bin", "--abi", "-o", outputPath, path)
	_, err := cmd.CombinedOutput()

	if err != nil {
		logger.Error("EthCompiler::CompileContract::execCommand:", err)
		return err
	}

	return nil
}
