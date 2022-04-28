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
func BorrowGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, borrowDays *big.Int,valueWei *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction,*big.Int,common.Address, error) {
	opts := GetMsgOpts(privateKey)
	opts.Value = valueWei
	res, err := contract.BorrowGoods(opts, id, borrowDays)

	fmt.Println("bGoods:", res)
	borrow,err:=contract.GoodsData(nil,id)
	fmt.Println("borrower",borrow.Borrowers.Borrower,borrowDays)

	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res,borrow.Deal,borrow.Borrowers.Borrower, nil
}

//封装同意借出方法
func AgreeMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, deal *big.Int, since string) (*types.Transaction, error)  {
	opts := Getopts()
	opts.Value = big.NewInt(1000000000000000000)
	res, err := contract.AgreeBorrow(opts,id,deal,since)

	fmt.Println("agree:", res)
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

