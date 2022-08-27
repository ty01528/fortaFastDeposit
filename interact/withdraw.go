package interact

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// Withdraw
//
//	@Description: 就是一个简单的mint过程
//	@Author: tianyuan01528@foxmail.com
//	@param toAddress
//	@param amount
func Withdraw(spenderAddress *big.Int) {
	// 连接rpc客户端
	client, err := ethclient.Dial(rpcClient)
	if err != nil {
		log.Fatalf("链接到RPC客户端失败! err: %v", err)
	}
	defer client.Close()

	// 读取私钥
	privateKey, err := crypto.HexToECDSA(contractOwnerPriKey)
	if err != nil {
		log.Fatalf("解析私钥失败! err: %v", err)
	}
	// 获取到私钥对应的公钥
	publicKey := privateKey.Public()
	// 签名事务
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 查询nonce与gas的价格
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("从RPC客户端查询nonce失败! err: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("从RPC客户端查询gas失败! err: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(int64(chainId)))
	if err != nil {
		log.Fatalf("创建Transactor失败! err: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 开始调用合约
	address := common.HexToAddress(contractAddress)
	act, err := NewInteractTransactor(address, client)
	if err != nil {
		log.Fatalf("合约调用异常! err: %v", err)
	}
	tx, err := act.Withdraw(auth, uint8(0), spenderAddress)
	if err != nil {
		println(spenderAddress, " 地址的’withdraw‘合约失败。")
	} else {
		println("账户:", spenderAddress.String(), "质押提取成功，交易哈希为：", tx.Hash().String())
	}
}
