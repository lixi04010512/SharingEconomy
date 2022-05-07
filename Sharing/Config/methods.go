package config

import (
	"Sharing/Agreement"
	"context"
	"crypto/ecdsa"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	//Prikey       = "769f41a9de1f6f1848d7631e274f2168e3cb20d16c81f688a49d081965a53861"
	//ShareFishAddress = "0x097Dbbe9AC7dDd12919A1C6C1D364852F5bcefdE"
	//hch
	//Prikey           = "5fd7eb82fbd4c5cc87fbeb867ff4804537723d80a5289606c28ca5c8eed4aa77"
	//ShareFishAddress = "0x271c8691cBbC597a13D560d197a45dBf5F5B070b"

	//lixi
	Prikey       = "45ee84c9fbfee72f88392d5e99fd071bfb1a89c2c1482456547e049d437c1380"
	ShareFishAddress = "0xA4300503eC7cC09ed4734F6d3724510530C7600a"

	//Prikey       = "45ee84c9fbfee72f88392d5e99fd071bfb1a89c2c1482456547e049d437c1380"
	//ShareFishAddress = "0xA4300503eC7cC09ed4734F6d3724510530C7600a"
	//Prikey       = "45ee84c9fbfee72f88392d5e99fd071bfb1a89c2c1482456547e049d437c1380"
	//ShareFishAddress = "0x72B9CE9e08c7e0C4d5698186cc5271F6dc382c1f"

	//Prikey       = "769f41a9de1f6f1848d7631e274f2168e3cb20d16c81f688a49d081965a53861"
	//ShareFishAddress = "0xDAafB847bE536849D698CC6497c9c1746Ad61be4"
	//Prikey       = "40fb59f1551424cde70fdc153b7e98cad6528e0680a1e7a3a400d0e2c3d27019"
	//ShareFishAddress = "0x63521Ca5ea88345C7AaF63F01dB793CA92eEceB0"

	gasLimit = 3000000
	//fileKeystore     = "UTC--2022-03-17T08-08-40.600466800Z--59b0f8a34d8f0dd0e0eef44d02cef0c12fffb9de"
	//Prikey = "c5b9c7fd467335bd829b3b2a3098a72ac39b7f5efa162220b7907cfc684df9a3"
	//privateKey       = "111"
	//ShareFishAddress = "0x7D1672baDeA4825c4B8fFD938fFD29457D6aF6ae"
)

//获取client
func GetClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	//client, err := ethclient.Dial("https://rinkeby.etherscan.io")

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	return client, nil
}
func HaveClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chainID)
	return client, err
}

//只写实例 获取UserCaller对象
func HaveUserWrite(client *ethclient.Client) (*Agreement.UserTransactor, error) {
	ins, err := Agreement.NewUserTransactor(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return ins, err
}

// 获取合约对象
func GetAddress(client *ethclient.Client) (*Agreement.User, error) {
	contract, err := Agreement.NewUser(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("faucet:", contract)
	return contract, err
}

//获取Opts对象
func Getopts() *bind.TransactOpts {
	privateKey := GetPrivateKey()
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}

//只读实例 获取UserCaller对象
func HaveUserRead(client *ethclient.Client) (*Agreement.UserCaller, error) {
	contract, err := Agreement.NewUserCaller(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return contract, err
}

//获取私钥
func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(Prikey)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return big.NewInt(0), err
	} else {
		fmt.Println("gasPrice", gasPrice)
		return gasPrice, nil
	}

}

func GetMsgOpts(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}
