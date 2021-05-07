package model

import (
	"errors"
	"github.com/StupidTAO/DIDHub/contracts"
	"github.com/StupidTAO/DIDHub/utils"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

const (
	//投票合约
	VOTE_CONTRACT_ADDRESS = "0x727eddf194b0b2b36b870ad0be7e69fdb2b37e2e"
)


func deployVote() (error) {
	blockchain, err := ethclient.Dial(URL)
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

	//构造提案名称
	var name [32]byte
	for i := 0; i < len(name); i++ {
		name[i] = '0'
	}
	name[31] = 'a'

	var name1[32] byte
	for i := 0; i < len(name1); i++ {
		name1[i] = '0'
	}
	name1[31] = 'b'

	//构造提案申请资金
	var funds []*big.Int
	funds = append(funds, big.NewInt(1000))
	funds = append(funds, big.NewInt(2000))
	var proposalNames [][32]byte
	proposalNames = append(proposalNames, name)
	proposalNames = append(proposalNames, name1)
	address, _, _, err := contracts.DeployVote(
		auth,
		blockchain,
		proposalNames,
		funds,
	)
	if err != nil {
		log.Fatalf("deploy %v \n", err)
		return err
	}

	fmt.Printf("Contract pending deploy:0x%x \n", address)
	return err
}

func ContractGiveRightToVote(voter common.Address, number int64) error {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(KEY), PASSWORD)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return err
	}

	tx, err := contract.GiveRightToVote(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, voter, big.NewInt(number))
	if err != nil {
		log.Fatalf("give right to voter: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractDelegate(to common.Address, prk string) error {
	//检查并获取私钥
	if len(prk) < 64 {
		return errors.New("private key content too short!")
	}

	pri, err := utils.GetPrivateKeyByString(prk)
	if err != nil {
		return err
	}

	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth := bind.NewKeyedTransactor(pri)

	tx, err := contract.Delegate(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, to)
	if err != nil {
		log.Fatalf("delegate ticket right failed: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractRevokeVote(prk string) error {
	//检查并获取私钥
	if len(prk) < 64 {
		return errors.New("private key content too short!")
	}

	pri, err := utils.GetPrivateKeyByString(prk)
	if err != nil {
		return err
	}

	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth := bind.NewKeyedTransactor(pri)

	tx, err := contract.RevokeVote(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	})
	if err != nil {
		log.Fatalf("revoke ticket failed: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractVote(prk string, proposal int64) error {
	//检查并获取私钥
	if len(prk) < 64 {
		return errors.New("private key content too short!")
	}

	pri, err := utils.GetPrivateKeyByString(prk)
	if err != nil {
		return err
	}

	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return err
	}

	auth := bind.NewKeyedTransactor(pri)

	tx, err := contract.Vote(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, big.NewInt(proposal))
	if err != nil {
		log.Fatalf("vote ticket failed: %v \n", err)
		return err
	}
	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return nil
}

func ContractWinningProposal() (int64, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return 65535, err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return 65535, err
	}

	winningProposal_, err := contract.WinningProposal(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("get winning propaosal failed: %v \n", err)
		return 65535, err
	}
	fmt.Printf("tx sent: %d \n", winningProposal_.Int64())
	return winningProposal_.Int64(), nil
}

func ContractWinnerName() (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	winnerName, err := contract.WinnerName(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("get winner name failed: %v \n", err)
		return "", err
	}

	return string(winnerName[:]), nil
}

func ContractChairperson() (string, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", err
	}

	chairAddr, err := contract.Chairperson(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("get chair person: %v \n", err)
		return "", err
	}

	return chairAddr.Hex(), nil
}

func ContractProposals(index int64) (string, int64, int64, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return "", 0, 0, err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return "", 0, 0, err
	}

	proposal, err := contract.Proposals(&bind.CallOpts{Pending: true}, big.NewInt(index))
	if err != nil {
		log.Fatalf("get propaosal failed: %v \n", err)
		return "", 0, 0, err
	}

	return string(proposal.Name[:]), proposal.VoteCount.Int64(), proposal.Funds.Int64(), nil
}

func ContractVoters(addr common.Address) (int64, bool, common.Address, int64, error) {
	blockchain, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v \n", err)
		return 0, false, common.Address{}, 0, err
	}

	contract, err := contracts.NewVote(common.HexToAddress(VOTE_CONTRACT_ADDRESS), blockchain)
	if err != nil {
		log.Fatalf("conn contract: %v \n", err)
		return 0, false, common.Address{}, 0, err
	}

	voter, err := contract.Voters(&bind.CallOpts{Pending: true}, addr)
	if err != nil {
		log.Fatalf("get voter failed: %v \n", err)
		return 0, false, common.Address{}, 0, err
	}

	return voter.Weight.Int64(), voter.Voted, voter.Delegate, voter.Vote.Int64(), nil
}
