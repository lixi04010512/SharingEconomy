package config

import (
	"Sharing/Agreement"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//封装捐赠积分方法
func DonateMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	opts := Getopts()
	opts.Value = big.NewInt(1000000000000000000)
	res, err := contract.Donate(opts, number, donator, amount, donatorName)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
