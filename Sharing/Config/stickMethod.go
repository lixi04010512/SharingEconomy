package config

import (
	"Sharing/Agreement"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//封装添加分类方法
func AddStick(client *ethclient.Client, contract *Agreement.User, stick string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddSticker(opts, stick)
	fmt.Println("addstick:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装分类展示方法
func ShowSpecies(contract *Agreement.User, id *big.Int) (string, error) {

	res, err := contract.GetStick(nil, id)

	//fmt.Println("showSpecies:", res)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装分类展示方法
func ShowStick(client *ethclient.Client, id *big.Int) (string, error) {
	ins, err := HaveUserRead(client)
	res, err := ins.GetStick(nil, id)

	//fmt.Println("showSpecies:", res)
	if err != nil {
		log.Fatal(err)
	}
	return res, err
}

//封装删除分类方法
func DelSpecieMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DelStick(opts, number)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	fmt.Println("Delstick:", err)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装修改分类方法
func UpdateStickMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int, name string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.UpdateStick(opts, number, name)
	fmt.Println("updateStick:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
