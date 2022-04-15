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

//封装物品上架方法
func AddGoodsMethod(client *ethclient.Client, contract *Agreement.User, owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string, goodsSign string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddGoods(opts, owner, name, species, rent, ethPledge, goodsImgs, goodsSign)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
