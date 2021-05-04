package model

import (
	"fmt"
	"github.com/StupidTAO/DIDHub/utils"
)

func InsertHubDIDDoc(didDoc DBDIDDoc) error {
	//插入数据库
	err := InsertDBDIDDoc(didDoc)
	if err != nil {
		fmt.Println("insert db did documnet error is ", err)
		return err
	}

	//插入区块链
	docsBytes := utils.GetSHA256HashCode([]byte(didDoc.DidDoc))
	docs58 := utils.Base58Encode(docsBytes)
	err = ContractSetDidDocumentHash(didDoc.Did, docs58)
	if err != nil {
		fmt.Println("insert blockchain did documnet hahs error is ", err)
		return err
	}

	fmt.Println("did doc insert did hub sucess")
	return nil
}

func FindHubDIDDoc(did string) ([]DBDIDDoc, error) {
	docs, err := FindDBDIDDoc(did)
	if err != nil {
		fmt.Println("find db did documnet error is ", err)
		return nil, err
	}

	if len(docs) == 0 {
		return nil, nil
	}

	docHash, err := ContractFindDidDocumentHash(did)
	if err != nil {
		fmt.Println("find blockchain did documnet error is ", err)
		return nil, err
	}

	docsBytes := utils.GetSHA256HashCode([]byte(docs[0].DidDoc))
	docs58 := utils.Base58Encode(docsBytes)
	if docs58 != docHash {
		fmt.Println("did document db and block chain dose not match")
		return nil, err
	}

	return docs, nil
}

func InsertHubDIDClaim(didClaim DBDIDClaim) error {
	//插入数据库
	err := InsertDBDIDClaim(didClaim)
	if err != nil {
		fmt.Println("insert db did claim error is ", err)
		return err
	}

	//插入区块链
	claimBytes := utils.GetSHA256HashCode([]byte(didClaim.DidClaim))
	claim58 := utils.Base58Encode(claimBytes)
	err = ContractSetDidClaimHash(didClaim.ClaimId, claim58)
	if err != nil {
		fmt.Println("insert blockchain did claim hahs error is ", err)
		return err
	}

	fmt.Println("did doc insert did claim sucess")
	return nil
}

func FindHubDIDClaim(claimId string) ([]DBDIDClaim, error) {
	claims, err := FindDBDIDClaim(claimId)
	if err != nil {
		fmt.Println("find db did claim error is ", err)
		return nil, err
	}

	if len(claims) == 0 {
		return nil, nil
	}

	claimHash, err := ContractFindDidClaimHash(claimId)
	if err != nil {
		fmt.Println("find blockchain did documnet error is ", err)
		return nil, err
	}

	claimBytes := utils.GetSHA256HashCode([]byte(claims[0].DidClaim))
	claim58 := utils.Base58Encode(claimBytes)
	if claim58 != claimHash {
		fmt.Println("did document db and block chain dose not match")
		return nil, err
	}

	return claims, nil
}

func InsertHubDIDPublicKey(didPublicKey DBDIDPublicKey) error {
	//插入数据库
	err := InsertDBDIDPublicKey(didPublicKey)
	if err != nil {
		fmt.Println("insert db did public key  error is ", err)
		return err
	}

	//插入区块链
	err = ContractSetDidPublicKey(didPublicKey.Did, didPublicKey.DidPublicKey)
	if err != nil {
		fmt.Println("insert blockchain did public key error is ", err)
		return err
	}

	fmt.Println("did doc insert did public key  sucess")
	return nil
}

func FindHubDIDPublicKey(did string) ([]DBDIDPublicKey, error) {
	//读取数据库中的内容
	publicKeys, err := FindDBDIDPublicKey(did)
	if err != nil {
		fmt.Println("find db did public key error is ", err)
		return nil, err
	}

	if len(publicKeys) == 0 {
		return nil, nil
	}

	//读取blockchain内容
	publicKey, err := ContractFindDidPublicKey(did)
	if err != nil {
		fmt.Println("find blockchain did public key error is ", err)
		return nil, err
	}

	if publicKeys[0].DidPublicKey != publicKey {
		fmt.Println("did public key db and block chain dose not match")
		return nil, err
	}

	return publicKeys, nil
}

func InsertHubDIDChainAddr(didChainAddr DBDIDChainAddr) error {
	//插入数据库
	err := InsertDBDIDChainAddr(didChainAddr)
	if err != nil {
		fmt.Println("insert db did chain addr error is ", err)
		return err
	}

	//插入区块链
	err = ContractSetDidChainAddr(didChainAddr.Did, didChainAddr.DidChainAddr)
	if err != nil {
		fmt.Println("insert blockchain did chain addr error is ", err)
		return err
	}

	fmt.Println("did doc insert did chain addr sucess")
	return nil
}

func FindHubDIDChainAddr(did string) ([]DBDIDChainAddr, error) {
	//读取数据库中的内容
	chainAddrs, err := FindDBDIDChainAddr(did)
	if err != nil {
		fmt.Println("find db did chain addr error is ", err)
		return nil, err
	}

	if len(chainAddrs) == 0 {
		return nil, nil
	}

	//读取blockchain内容
	chainAddr, err := ContractFindDidChainAddr(did)
	if err != nil {
		fmt.Println("find blockchain did public key error is ", err)
		return nil, err
	}

	if chainAddrs[0].DidChainAddr != chainAddr {
		fmt.Println("did public key db and block chain dose not match")
		return nil, err
	}

	return chainAddrs, nil
}

func InsertHubTransaction(tx Transaction) error {
	//插入数据库
	err := InsertDBTransaction(&tx)
	if err != nil {
		fmt.Println("insert db transaction error is ", err)
		return err
	}

	//插入区块链
	err = ContractSetTransaction(tx.TxId, tx.FromAddr, tx.ToAddr, int64(tx.Amount))
	if err != nil {
		fmt.Println("insert blockchain transaction error is ", err)
		return err
	}

	fmt.Println("insert hub transaction sucess")
	return nil
}

func FindHubTranaction(txId string) ([]Transaction, error) {
	//读取数据库中的内容
	txs, err := FindDBTransaction(txId)
	if err != nil {
		fmt.Println("find db transaction error is ", err)
		return nil, err
	}

	if len(txs) == 0 {
		return nil, nil
	}

	//读取blockchain内容
	fromAddr, toAddr, amount, err := ContractFindTransaction(txId)
	if err != nil {
		fmt.Println("find blockchain transaction error is ", err)
		return nil, err
	}

	if txs[0].FromAddr != fromAddr || txs[0].ToAddr != toAddr || int64(txs[0].Amount) != amount{
		fmt.Println("trnasaction db and block chain dose not match")
		return nil, err
	}

	return txs, nil
}
