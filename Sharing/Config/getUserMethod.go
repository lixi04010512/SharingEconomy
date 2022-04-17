package config

import (
	"Sharing/Agreement"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//封装获取用户信息方法
func GetUserMethod(contract *Agreement.User, address common.Address) (string, common.Address, *big.Int, string, string, *big.Int, *big.Int, error) {
	res, err := contract.GetUser(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return res.Name, res.People, res.Integral, res.Email, res.Sign, res.GoodsNum, res.Balance, nil
}

func UpdateUserMethod(client *ethclient.Client,contract *Agreement.User, people common.Address, name string, email string, sign string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error)  {
	opts := GetMsgOpts(privateKey)
	res, err := contract.UpdateUser(opts,people,name,email,sign)
	fmt.Println("up:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
