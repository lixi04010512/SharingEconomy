package config

import (
	"Sharing/Agreement"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
)

//封装充值方法
func TopUpMethod(client *ethclient.Client, contract *Agreement.User,myAmount *big.Int ,people common.Address) (*types.Transaction, error)  {
	opts := Getopts()

	num, err := strconv.Atoi(myAmount.String())
	opts.Value = big.NewInt(int64(num)*1000000000000000000)
	res, err := contract.TopUp(opts,myAmount,people)

	fmt.Println("topUp:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
