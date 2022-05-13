package config

import (
	"Sharing/Agreement"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

//封装登录方法
func LoginMethod(client *ethclient.Client, contract *Agreement.User, Address common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.Login(opts, Address)

	fmt.Println("login:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil

}
