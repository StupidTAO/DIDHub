package model

import (
	"fmt"
	"testing"
)

const keyTest=`{"address":"12769c3419a7f491cf4e576e2e983e009d579076","crypto":{"cipher":"aes-128-ctr","ciphertext":"215430a18ab1132c6eaecdf966bc0d878a3be06cff5dce173d801afec5002db5","cipherparams":{"iv":"d41d87954da3dfca1f38e14111169fb8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d5268e70fbf8666435bf82ee53850f14486810f1944110de2aede933ae97fff1"},"mac":"a80e95fc657473f543bb989da9e8e2cde73e2b8bde52e6b05355b44a40c165ec"},"id":"5e034459-6e81-4208-9e26-c91641d20f5d","version":3}`


func TestDeploy(t *testing.T) {
	//Contract pending deploy
	deploy()

}

func TestContractSetDidChainAddr(t *testing.T) {
	err := ContractSetDidChainAddr("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "0x5a14c2136084A24853EfF907185f79AE77c36c7c")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractFindDidChainAddr(t *testing.T) {
	ethAddress, err := ContractFindDidChainAddr("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(ethAddress)
}


func TestContractSetDidClaimHash(t *testing.T) {
	err := ContractSetDidClaimHash("0xclaim_id12345678", "0xclaimHash1234567890")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractFindDidClaimHash(t *testing.T) {
	claimHash, err := ContractFindDidClaimHash("0xclaim_id12345678")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(claimHash)
}

func TestContractSetDidDocumentHash(t *testing.T) {
	err := ContractSetDidDocumentHash("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "0xdocumentHash1234567890")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractFindDidDocumentHash(t *testing.T) {
	claimHash, err := ContractFindDidDocumentHash("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(claimHash)
}

func TestContractSetDidPublicKey(t *testing.T) {
	err := ContractSetDidPublicKey("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "0xpublicKey1234567890")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractFindDidPublicKey(t *testing.T) {
	publicKey, err := ContractFindDidPublicKey("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(publicKey)
}

func TestContractSetTransaction(t *testing.T) {
	err := ContractSetTransaction("10001","did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1", 123)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestContractFindTransaction(t *testing.T) {
	from, to, amount, err := ContractFindTransaction("10001")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)
}
