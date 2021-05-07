package model

import (
	"fmt"
	"github.com/StupidTAO/DIDHub/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

const (
	KEY = `{"address":"12769c3419a7f491cf4e576e2e983e009d579076","crypto":{"cipher":"aes-128-ctr","ciphertext":"215430a18ab1132c6eaecdf966bc0d878a3be06cff5dce173d801afec5002db5","cipherparams":{"iv":"d41d87954da3dfca1f38e14111169fb8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d5268e70fbf8666435bf82ee53850f14486810f1944110de2aede933ae97fff1"},"mac":"a80e95fc657473f543bb989da9e8e2cde73e2b8bde52e6b05355b44a40c165ec"},"id":"5e034459-6e81-4208-9e26-c91641d20f5d","version":3}`
	PASSWORD = "abc"
	URL = "http://127.0.0.1:7545"
	CONTRACT_ADDRESS = "0xa252df5c19e6b2bbfd1308b6c0d8e0430a134ce3"
)

func deploy() (error) {
	blockchain, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	// 合约部署
	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	fmt.Printf("address: %s\n", auth.From.String())
	address, _, _, err := contracts.DeployDataStorage(
		auth,
		blockchain,
	)
	if err != nil {
		log.Fatalf("deploy %v \n", err)
		return err
	}

	fmt.Printf("Contract pending deploy:0x%x \n", address)

	return err
}

//设置DID与对应的以太坊地址, "0xf95ca40fd9bfa2c36728b6538865c5a6f3f8d7bd"
//将这三个参数写入配置文件中，url、kyefile、address
func ContractSetDidChainAddr(did, _didChainAddr string) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.SetDidChainAddr(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, did, _didChainAddr)
	if err != nil {
		log.Fatalf("set did chain addr: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractFindDidChainAddr(did string) (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	ethAddress, err := contract.GetDidChainAddr(&bind.CallOpts{Pending: true}, did)
	if err != nil {
		log.Fatalf("Failed to get did ethAddress:%v \n", err)
		return "", err
	}

	return ethAddress, nil
}

func ContractSetDidClaimHash(claimId, _didClaimHash string) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.SetDidClaimHash(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, claimId, _didClaimHash)
	if err != nil {
		log.Fatalf("set did claim: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractFindDidClaimHash(claimId string) (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	claimHash, err := contract.GetDidClaimHash(&bind.CallOpts{Pending: true}, claimId)
	if err != nil {
		log.Fatalf("Failed to get did claim :%v \n", err)
		return "", err
	}

	return claimHash, nil
}

func ContractSetDidDocumentHash(did, _didDocumentHash string) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.SetDidDocumentHash(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, did, _didDocumentHash)
	if err != nil {
		log.Fatalf("set did document: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractFindDidDocumentHash(did string) (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	documentHash, err := contract.GetDidDocumentHash(&bind.CallOpts{Pending: true}, did)
	if err != nil {
		log.Fatalf("Failed to get did ethAddress:%v \n", err)
		return "", err
	}

	return documentHash, nil
}

func ContractSetDidPublicKey(did, _didPublicKey string) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.SetDidPublicKey(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, did, _didPublicKey)
	if err != nil {
		log.Fatalf("set did public key: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractFindDidPublicKey(did string) (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	publicKey, err := contract.GetDidPublicKey(&bind.CallOpts{Pending: true}, did)
	if err != nil {
		log.Fatalf("Failed to get did ethAddress:%v \n", err)
		return "", err
	}

	return publicKey, nil
}

//向区块链写入交易
func ContractSetTransaction(txId string, fromAddr string, toAddr string, amount int64) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.SetTransaction(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, txId, fromAddr, toAddr, big.NewInt(amount))

	if err != nil {
		log.Fatalf("set transaciton err: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractFindTransaction(txId string) (string, string, int64, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", "", 0, err
	}

	contract, err := contracts.NewDataStorage(common.HexToAddress(CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", "", 0, err
	}

	fromAddr, toAddr, amount, err := contract.GetTransaction(&bind.CallOpts{Pending: true}, txId)
	if err != nil {
		log.Fatalf("Failed to get did ethAddress:%v \n", err)
		return "", "", 0, err
	}

	return fromAddr, toAddr, amount.Int64(), nil
}
