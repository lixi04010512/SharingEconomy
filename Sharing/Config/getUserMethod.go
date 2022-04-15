package config

import (
	"Sharing/Agreement"
	"github.com/ethereum/go-ethereum/common"
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
