package config

import (
	"Sharing/Agreement"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//封装物品下架方法
func DelGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DelGoods(opts, id)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品借出方法
func BorrowGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.BorrowGoods(opts, id, borrowDays)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品归还方法
func DoGoodsReturnMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DoGoodsReturn(opts, id)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
