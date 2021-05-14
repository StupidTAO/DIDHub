package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

const (
	PRIVATE_KEY_TEST = "a58c86588a45812a5e9fbb779b3993f3c1b220f213456d7c8a856892dd60f3cf"
	ADDRESS_TEST = "0x087b8951a4161bAB2f474bE3fD28F9154a221D45"
	PRIVATE_KEY_TEST1 = "237667a226cb4d605490a220ca18371f19e4b4a2efcf1a143a7a1f2c4d1db5ba"
	ADDRESS_TEST1 = "0xDd6018c423634c59CD7b90146D51eC4e97Ac3D1d"
	CHAIR_ADDRESS = "0x12769c3419a7f491cf4e576e2e983e009d579076"
)


func TestDeployVote(t *testing.T) {
	//Contract pending deploy
	deployVote()
}

func TestContractGiveRightToVote(t *testing.T) {
	err := ContractGiveRightToVote(common.HexToAddress(ADDRESS_TEST), 100)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractDelegate(t *testing.T) {
	err := ContractDelegate(common.HexToAddress(ADDRESS_TEST1), PRIVATE_KEY_TEST)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractRevokeVote(t *testing.T) {
	err := ContractRevokeVote(PRIVATE_KEY_TEST)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractVote(t *testing.T) {
	err := ContractVote(PRIVATE_KEY_TEST1,1)
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
	name, count, needFunds, err := ContractProposals(1)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("proposal name is: %s, voted count is: %d, proposal need funds is: %d\n", name, count, needFunds)
}

func TestContractVoters(t *testing.T) {
	weight, voted, addr, vote, err := ContractVoters(common.HexToAddress("0x21BEd1452BAf5D5a952dde3F2f8Eb6c3a5be798b"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("voter weight is: %d, voted is: %t, delegate is: %s, vote index is: %d\n", weight, voted, addr.Hex(), vote)
}
