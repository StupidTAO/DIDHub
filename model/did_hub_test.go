package model

import (
	"fmt"
	"github.com/StupidTAO/DIDHub/utils"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestInsertHubDIDDoc(t *testing.T) {
	doc := new(DBDIDDoc)
	doc.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	doc.DidDoc = "{\"Context\":\"https://w3id.org/did/v1\",\"PublicKey\":{\"Id\":\"#keys-1\",\"Type\":\"Secp256k1\",\"PublicKeyHex\":\"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71\"},\"Authentication\":\"#key-1\"}\n"
	doc.IsAvailable = 1
	doc.CreateTime = time.Now().Add(8 * time.Hour)
	doc.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertHubDIDDoc(*doc)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindHubDIDDoc (t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	docs, err := FindHubDIDDoc(did)
	if err != nil {
		t.Error("error is ", err)
	}
	fmt.Println(len(docs))
}

func TestInsertHubDIDClaim(t *testing.T) {
	claim := new(DBDIDClaim)
	claim.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	claim.DidClaim = "{\"Id\":\"4TDE7qofZiKAwfbxujbnbwehkBpV\",\"Issuer\":\"did:welfare:3rwXEnN4PRQ11QBiDbtGQYBGwGea\",\"IssuanceDate\":\"2021-04-13 00:19:32.041385 +0800 CST m=+60272.449758779\",\"ExpirationDate\":\"2021-04-13 00:19:32.041387 +0800 CST m=+60272.449761052\",\"Result\":1,\"CredentialSubject\":{\"id\":\"did:welfare:3L9EUK88XxMiNu9SnYPEGTFFUGzV\",\"shortDescription\":\"342225199509082342\",\"longDescription\":\"ID Card\",\"typeClaim\":\"IDCardAuthentication\"},\"Proof\":{\"creator\":\"did:welfare:3rwXEnN4PRQ11QBiDbtGQYBGwGea\",\"type\":\"Keccak256\",\"signatureValue\":\"HEHUrhQ9UGdKRcpgqwfqBvf9xQxwYV3BL7hjJPmFbUT96Tf5w19jCQQso8E3bedWDYGk8CS3teXEhjfdMoYY6wHNY\"}}"
	bs := utils.GetRipemd160HashCode([]byte(claim.DidClaim))
	claim.ClaimId = utils.Base58Encode(bs)
	claim.IsAvailable = 1
	claim.CreateTime = time.Now().Add(8 * time.Hour)
	claim.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertHubDIDClaim(*claim)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindHubDIDClaim (t *testing.T) {
	claimId := "2pt76zHa7A8EtiDjyGY1P7DYGNVN"

	claims, err := FindHubDIDClaim(claimId)
	if err != nil {
		t.Error("error is ", err)
	}
	fmt.Println(len(claims))
}

func TestInsertHubDIDPublicKey(t *testing.T) {
	publicKey := new(DBDIDPublicKey)
	publicKey.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	publicKey.DidPublicKey = "215430a18ab1132c6eaecdf966bc0d878a3be06cff5dce173d801afec5002db5"
	publicKey.IsAvailable = 1
	publicKey.CreateTime = time.Now().Add(8 * time.Hour)
	publicKey.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertHubDIDPublicKey(*publicKey)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindHubDIDPublicKey (t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	claims, err := FindHubDIDPublicKey(did)
	if err != nil {
		t.Error("error is ", err)
	}
	fmt.Println(len(claims))
}

func TestInsertHubDIDChainAddr(t *testing.T) {
	chainAddr := new(DBDIDChainAddr)
	chainAddr.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	chainAddr.DidChainAddr = "0x5a14c2136084A24853EfF907185f79AE77c36c7c"
	chainAddr.IsAvailable = 1
	chainAddr.CreateTime = time.Now().Add(8 * time.Hour)
	chainAddr.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertHubDIDChainAddr(*chainAddr)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindHubDIDChainAddr (t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	chainAddrs, err := FindHubDIDChainAddr(did)
	if err != nil {
		t.Error("error is ", err)
	}
	fmt.Println(len(chainAddrs))
}

func TestInsertHubTransaction(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	amount := uint32(rand.Intn(1000))
	tx := new(Transaction)
	tx.TxId = strconv.Itoa(10001)
	tx.FromAddr = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	tx.ToAddr = "did:welfare:123456789abcdefghijklmnopqrs"
	tx.Amount = amount
	tx.ProjectPriority = 1.4
	tx.Contribution = 0
	tx.CreateTime = time.Now().Add((24*14+8) * time.Hour)
	tx.UpdateTime = time.Now().Add((24*14+8) * time.Hour)
	tx.HasCaculate = 0
	tx.Payload = "test-0"

	err := InsertHubTransaction(*tx)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindHubTransaciton (t *testing.T) {
	txId := "10001"
	txs, err := FindHubTranaction(txId)
	if err != nil {
		t.Error("error is ", err)
	}
	fmt.Println(len(txs))
}
