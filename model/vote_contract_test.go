package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestDeployVote(t *testing.T) {
	//Contract pending deploy
	deployVote()
}

func TestContractGiveRightToVote(t *testing.T) {
	err := ContractGiveRightToVote(common.HexToAddress("0x087b8951a4161bAB2f474bE3fD28F9154a221D"), 100)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractDelegate(t *testing.T) {
	err := ContractDelegate(common.HexToAddress("0x087b8951a4161bAB2f474bE3fD28F9154a221D"))
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractRevokeVote(t *testing.T) {
	err := ContractRevokeVote()
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractVote(t *testing.T) {
	err := ContractVote(1)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractWinningProposal(t *testing.T) {
	proposal, err := ContractWinningProposal()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("proposal id is: %d\n", proposal)
}

func TestContractWinnerName(t *testing.T) {
	winnerName, err := ContractWinnerName()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("proposal winner name is: %s\n", winnerName)
}

func TestContractChairperson(t *testing.T) {
	chairPerson, err := ContractChairperson()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("proposal winner name is: %s\n", chairPerson)
}

func TestContractProposals(t *testing.T) {
	name, count, err := ContractProposals(1)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("proposal name is: %s, voted count is: %d\n", name, count)
}

func TestContractVoters(t *testing.T) {
	weight, voted, addr, vote, err := ContractVoters(common.HexToAddress("0x12769c3419A7f491CF4e576e2E983e009d579076"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("voter weight is: %d, voted is: %t, delegate is: %s, vote count is: %d\n", weight, voted, addr.Hex(), vote)
}
