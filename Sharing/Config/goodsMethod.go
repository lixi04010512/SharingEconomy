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

//封装物品下架方法
func OutGoodsMethod(client *ethclient.Client, contract *Agreement.User, owner common.Address,id *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error)  {
	opts := GetMsgOpts(privateKey)
	res, err := contract.OutGoods(opts,owner,id)

	fmt.Println("del:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}


//修改物品
//func UpdateGoodsMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, name string, species string, rent *big.Int, ethPledge *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error)   {
//	opts := GetMsgOpts(privateKey)
//	res, err := contract.UpdGoods(opts,id,name,species,rent,ethPledge)
//	fmt.Println("update:", res)
//	opts.GasLimit = gasLimit
//	opts.GasPrice, err = GetgasPrice(client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return res, nil
//}