# DIDHub

## 操作步骤
1.在db/connect.go文件InitDB函数中填写mysql数据库配置信息，如user、password、dbname

2.首先在本地运行ganache，设置rpc server为HTTP://127.0.0.1:7545

3.运行model/did_contract_test.go中的TestDeploy函数，部署智能合约（为区块链上hub存储合约）

4.将部署后的合约地址，填写到model/did_contract.go文件CONTRACT_ADDRESS常量中

5.运行model/vote_contract_test.go中的TestDeployVote函数，部署投票合约（为区块链上流动民主投票合约）

6.将部署后的合约地址，填写到model/vote_contract.go文件VOTE_CONTRACT_ADDRESS常量中

准备好智能合约后，即可以被其它模块调用，数据将存储在mysql和智能合约中。

## 模块功能

本模块为DIDHub模块，为DID Wallet的核心模块，该模块可以实现将DID相关信息存储在mysql中，并上传至区块链。区块链负责存储DID的相关信息，区块链选增以太坊。
