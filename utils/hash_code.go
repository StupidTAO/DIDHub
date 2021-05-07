package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

func GetSHA256HashCode(message []byte) []byte {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//返回哈希值
	return bytes
}

func GetRipemd160HashCode(message []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(message)
	hashBytes := hasher.Sum(nil)
	return hashBytes
}

func Base58Encode(message []byte) string {
	return base58.Encode(message)
}

func Base58Decode(msgStr string) []byte {
	return base58.Decode(msgStr)
}

func GetPrivateKeyByString(prk string) (*ecdsa.PrivateKey, error) {
	if len(prk) < 64 {
		return nil, errors.New("private file content too short")
	}
	//省去最后一个换行符
	ecdsaKey, err := crypto.HexToECDSA(prk[:64])
	if err != nil {
		return nil, err
	}
	return ecdsaKey, nil
}
