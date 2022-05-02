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
//封装物品借出方法
func AgreeBorrowMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, borrowDays *big.Int,valueWei *big.Int, since string) (*types.Transaction, error) {
	opts := Getopts()
	opts.Value = valueWei
	res, err := contract.AgreeBorrow(opts, id, borrowDays,since)

	fmt.Println("AGoods:", res)

	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装借出方法
func BorrowMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, deal *big.Int) (*types.Transaction, error)  {
	opts := Getopts()
	res, err := contract.Borrow(opts,id,deal)

	fmt.Println("Borrow:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
//封装不同意借出方法
func DisagreeMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, deal *big.Int,value *big.Int) (*types.Transaction, error)  {
	opts := Getopts()
	opts.Value = value
	res, err := contract.DisagreeBorrow(opts,id,deal)

	fmt.Println("agree:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}


//封装物品归还方法
func DoGoodsReturnMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction,*big.Int, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.DoGoodsReturn(opts, id)
	fmt.Println("doGoods:", res)
	back,err:=contract.GoodsData(nil,id)
	backId:=back.Backs
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res,backId, nil
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

//封装不同意归还方法
func DisagreeBackMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, backs *big.Int,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
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
func BackGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, backs *big.Int, over string,value *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	opts.Value =value
	res, err := contract.BackGoods(opts, id,backs,over)
	fmt.Println("back:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装借出方法
func HashMethod(client *ethclient.Client, contract *Agreement.User,id *big.Int, deal *big.Int,blockNum *big.Int) (*types.Transaction, error)  {
	opts := Getopts()
	res,err := contract.DealHash(opts,id,deal,blockNum)

	fmt.Println("Borrow:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

