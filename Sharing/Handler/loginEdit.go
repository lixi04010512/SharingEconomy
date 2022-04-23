package handler

import (
	config "Sharing/Config"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)


//用户地址
var LoginUser common.Address
var privKey *ecdsa.PrivateKey

//登录
func login(c *gin.Context) {
	fmt.Println("login")
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	_, f, err := c.Request.FormFile("log_addr")
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	password := c.PostForm("log_passwd")

	keyjson, errs := ioutil.ReadFile("./keystore/" + f.Filename)
	fmt.Println("keyjson:", string(keyjson))
	if errs != nil {
		respError(c, err)
		return
	}
	unlockedKey, errors := keystore.DecryptKey(keyjson, password)
	fmt.Println("unlock:", unlockedKey)
	if errors != nil {
		respError(c, err)
		fmt.Println(errors)
		return
	}
	privKey = unlockedKey.PrivateKey
	comAddr := unlockedKey.Address
	LoginUser = comAddr
	data, err := config.LoginMethod(client, contract, comAddr, privKey)
	addr := comAddr.Hex()

	c.Redirect(http.StatusFound, "/index")
	//respOK(c,data)
	fmt.Println("addr", addr)
	fmt.Println("data", data)
}
//私钥登录
func privateLogin(c *gin.Context) {
	fmt.Println("ok")
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	privateKeyStr:= c.PostForm("log_privateKey")
	fmt.Println("str",privateKeyStr)
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}


	privKey, _ = crypto.HexToECDSA(privateKeyStr)
	fmt.Println("prikey",privKey)

	comAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	fmt.Println("comaddr",comAddr)
	LoginUser = comAddr
	data, err := config.LoginMethod(client, contract, comAddr, privKey)
	addr := comAddr.Hex()

	c.Redirect(http.StatusFound, "/index")
	//respOK(c,data)
	fmt.Println("addr", addr)
	fmt.Println("data", data)
}
