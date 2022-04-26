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
	//chainID = 8565 //8888
	//Prikey       = "769f41a9de1f6f1848d7631e274f2168e3cb20d16c81f688a49d081965a53861"
	//ShareFishAddress = "0x097Dbbe9AC7dDd12919A1C6C1D364852F5bcefdE"
	//hch
	//Prikey       = "5fd7eb82fbd4c5cc87fbeb867ff4804537723d80a5289606c28ca5c8eed4aa77"
	//ShareFishAddress = "0xE2FDDD21499fa8bc26ae16DA651cC5B66731c878"
	//Prikey           = "5fd7eb82fbd4c5cc87fbeb867ff4804537723d80a5289606c28ca5c8eed4aa77"
	//ShareFishAddress = "0xE7C0dFA71AEAbB67d73654c07eD4D79E66Bd53ce"

	//lixi
	//Prikey       = "45ee84c9fbfee72f88392d5e99fd071bfb1a89c2c1482456547e049d437c1380"
	//ShareFishAddress = "0xA714808022fd3840887149fe5E9B93ed52946bC7"
	//Prikey       = "45ee84c9fbfee72f88392d5e99fd071bfb1a89c2c1482456547e049d437c1380"
	//ShareFishAddress = "0x72B9CE9e08c7e0C4d5698186cc5271F6dc382c1f"
	Prikey       = "769f41a9de1f6f1848d7631e274f2168e3cb20d16c81f688a49d081965a53861"
	ShareFishAddress = "0x2632126fD88bf87810c05Cbd3f0dd4F843dBc628"

	gasLimit = 3000000
	//fileKeystore     = "UTC--2022-03-17T08-08-40.600466800Z--59b0f8a34d8f0dd0e0eef44d02cef0c12fffb9de"
	//Prikey = "c5b9c7fd467335bd829b3b2a3098a72ac39b7f5efa162220b7907cfc684df9a3"
	//privateKey       = "111"
	//ShareFishAddress = "0xAe03C34544973DEc93464121859fF8F7203d75fC"
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

//获取账户文件和密码 opts
//func HaveGetOpts() *bind.TransactOpts {
//	//账户文件
//	b, err := ioutil.ReadFile(fileKeystore)
//	if err != nil {
//		log.Fatal("错误", err)
//	}
//	//账户密码
//	key, err := keystore.DecryptKey(b, privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("key:", key)
//	reader, _ := os.Open(fileKeystore)
//	//opts, err := bind.NewTransactor(reader, "111")
//	opts, err := bind.NewTransactorWithChainID(reader, "111", big.NewInt(9696))
//	if err != nil {
//		log.Fatal("NewTransactor ", err)
//	}
//	return opts
//}

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
