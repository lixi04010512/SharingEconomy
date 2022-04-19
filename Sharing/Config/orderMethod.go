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


//封装物品借出方法
func BorrowGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, borrowDays *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.BorrowGoods(opts, id, borrowDays)
	fmt.Println("bGoods:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装同意借出方法
func AgreeMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, deal *big.Int, since string, borrower common.Address,privateKey *ecdsa.PrivateKey) (*types.Transaction, error)  {
	opts := GetMsgOpts(privateKey)
	res, err := contract.AgreeBorrow(opts,id,deal,since,borrower)

	fmt.Println("agree:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品借出方法
func BorrowMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, deal *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.Borrow(opts, id, deal)
	fmt.Println("borrow:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品归还方法
func DoGoodsReturnMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.DoGoodsReturn(opts, id)
	fmt.Println("doGoods:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
//封装同意归还方法
func AgreeBackMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, backs *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.AgreeBack(opts, id,backs)
	fmt.Println("aback:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装归还方法
func BackGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, backs *big.Int, over string,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.BackGoods(opts, id,backs,over)
	fmt.Println("back:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

