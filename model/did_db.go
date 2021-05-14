package model

import (
	"errors"
	"fmt"
	"github.com/StupidTAO/DIDHub/db"
	"time"
)

type DBDIDDoc struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidDoc string			`db:"did_doc"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDClaim struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	ClaimId string			`db:"claimId"`
	DidClaim string			`db:"didClaim"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDPublicKey struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidPublicKey string		`db:"did_public_key"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDChainAddr struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidChainAddr string		`db:"did_chain_addr"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type Transaction struct {
	Id uint					`db:"id"`
	TxId string				`db:"txId"`
	FromAddr string			`db:"fromAddr"`
	ToAddr string			`db:"toAddr"`
	Payload string			`db:"payload"`
	Amount uint32			`db:"amount"`
	ProjectPriority float32 `db:"projectPriority"`
	Contribution float32	`db:"contribution"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	HasCaculate uint		`db:"hasCaculate"`
}

func InsertDBDIDDoc(didDoc DBDIDDoc) error {
	db.InitDB()

	//先查询是否存在
	docs, _ := FindDBDIDDoc(didDoc.Did)
	if len(docs) != 0 {
		return errors.New("did doc alread exist")
	}
	sql := "insert into did_document(did, did_doc, createTime, updateTime, isAvailable)values (?,?,?,?,?)"

	//执行SQL语句
	_, err := db.DB.Exec(sql, didDoc.Did, didDoc.DidDoc, didDoc.CreateTime, didDoc.UpdateTime, didDoc.IsAvailable)
	if err != nil {
		log.Info("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDDoc(did string) ([]DBDIDDoc, error){
	DB := db.InitDB()

	var docs []DBDIDDoc
	sql := "select id, did, did_doc, createTime, updateTime, isAvailable from did_document where did=?"
	err := DB.Select(&docs, sql, did)
	if err != nil {
		log.Info("exec failed, ", err)
		return nil, err
	}
	return docs, nil
}

func InsertDBDIDClaim(didClaim DBDIDClaim) error {
	sql := "insert into did_claim(did, claimId, didClaim, createTime, updateTime, isAvailable)values (?,?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	//查询是否存在
	claims, _ := FindDBDIDClaim(didClaim.ClaimId)
	if len(claims) != 0 {
		return errors.New("did claim alread exist")
	}

	_, err := db.DB.Exec(sql, didClaim.Did, didClaim.ClaimId, didClaim.DidClaim, didClaim.CreateTime, didClaim.UpdateTime, didClaim.IsAvailable)
	if err != nil {
		log.Info("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDClaim(did string) ([]DBDIDClaim, error){
	DB := db.InitDB()

	var claims []DBDIDClaim
	sql := "select id, did, claimId, didClaim, createTime, updateTime, isAvailable from did_claim where claimId=?"
	err := DB.Select(&claims, sql, did)
	if err != nil {
		log.Info("exec failed, ", err)
		return nil, err
	}
	return claims, nil
}

func InsertDBDIDPublicKey(didPublicKey DBDIDPublicKey) error {
	sql := "insert into did_publickey(did, did_public_key, createTime, updateTime, isAvailable)values (?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	//查询是否存在
	publicKeys, _ := FindDBDIDPublicKey(didPublicKey.Did)
	if len(publicKeys) != 0 {
		return errors.New("did publicKey alread exist")
	}

	_, err := db.DB.Exec(sql, didPublicKey.Did, didPublicKey.DidPublicKey, didPublicKey.CreateTime, didPublicKey.UpdateTime, didPublicKey.IsAvailable)
	if err != nil {
		log.Info("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDPublicKey(did string) ([]DBDIDPublicKey, error){
	DB := db.InitDB()

	var publickeys []DBDIDPublicKey
	sql := "select id, did, did_public_key, createTime, updateTime, isAvailable from did_publickey where did=?"
	err := DB.Select(&publickeys, sql, did)
	if err != nil {
		log.Info("exec failed, ", err)
		return nil, err
	}
	return publickeys, nil
}

func InsertDBDIDChainAddr(didChainAddr DBDIDChainAddr) error {
	sql := "insert into did_chain_addr(did, did_chain_addr, createTime, updateTime, isAvailable)values (?,?,?,?,?)"
	//查询是否存在
	chaiAddrs, _ := FindDBDIDChainAddr(didChainAddr.Did)
	if len(chaiAddrs) != 0 {
		return errors.New("did chaiAddr alread exist")
	}

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, didChainAddr.Did, didChainAddr.DidChainAddr, didChainAddr.CreateTime, didChainAddr.UpdateTime, didChainAddr.IsAvailable)
	if err != nil {
		log.Info("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDChainAddr(did string) ([]DBDIDChainAddr, error){
	DB := db.InitDB()

	var chainAddrs []DBDIDChainAddr
	sql := "select id, did, did_chain_addr, createTime, updateTime, isAvailable from did_chain_addr where did=?"
	err := DB.Select(&chainAddrs, sql, did)
	if err != nil {
		log.Info("exec failed, ", err)
		return nil, err
	}
	return chainAddrs, nil
}

func InsertDBTransaction(tx *Transaction) error {
	sql := "insert into transactions(txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate)values (?,?,?,?,?,?,?,?,?,?)"

	//查询是否存在
	txs, _ := FindDBTransaction(tx.TxId)
	if len(txs) != 0 {
		return errors.New("transaction alread exist")
	}

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, tx.TxId, tx.FromAddr, tx.ToAddr, tx.Payload, tx.Amount, tx.ProjectPriority, tx.Contribution, tx.CreateTime, tx.UpdateTime, tx.HasCaculate)
	if err != nil {
		log.Info("exec failed,", err)
		return err
	}

	return nil
}

func FindDBTransaction(txId string) ([]Transaction, error){
	DB := db.InitDB()

	var txs []Transaction
	var sql = "select id, txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate from transactions where txId=?"
	err := DB.Select(&txs, sql, txId)
	if err != nil {
		log.Info("exec failed, ", err)
		return nil, err
	}
	return txs, nil
}

func GetWelfareToken(fromAddr string) (float32, error) {
	DB := db.InitDB()

	var txs []Transaction
	var sql = "select id, txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate from transactions where fromAddr=?"
	err := DB.Select(&txs, sql, fromAddr)
	if err != nil {
		log.Info("exec failed, ", err)
		return 0, err
	}
	var token float32
	for i := 0; i < len(txs); i++ {
		token += txs[i].Contribution
	}

	return token, nil
}
