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

//封装添加公益方法
func AddCommunityMethod(client *ethclient.Client, contract *Agreement.User, beneficiaryAddr common.Address, communityName string, communityIntroduce string, communityAmounts *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddCommunity(opts, beneficiaryAddr, communityName, communityIntroduce, communityAmounts)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装下架公益方法
func DelCommunityMethod(client *ethclient.Client,contract *Agreement.User, number *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res,err := contract.DelCommunity(opts, number)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	fmt.Println("DelCommunity:", err)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装公益展示方法
func ShowCommunities(contract *Agreement.User, id *big.Int) (string, common.Address, string, *big.Int,  error) {
	community, err := contract.GetCommunity(nil, id)
	fmt.Println("showCommunities:", community)
	if err != nil {
		log.Fatal(err)
	}
	return community.Name, community.People, community.Introduce, community.Amount, nil
}

//封装修改公益方法
func UpdateCommunityMethod(client *ethclient.Client, contract *Agreement.User,name string,number *big.Int,addr common.Address,introduce string,amount *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.UpdateCommunity(opts, number,addr, name,introduce,amount)
	fmt.Println("updateCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
