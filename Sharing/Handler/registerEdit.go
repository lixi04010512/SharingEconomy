package handler

import (
	config "Sharing/Config"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

//邮箱验证
func sendEmail(c *gin.Context) {
	email := c.PostForm("use_email")
	var code = config.EmailSend(email)
	respOK(c, code)
}

//注册
var names string
func register(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	name := c.PostForm("use_name")
	email := c.PostForm("use_email")
	password := c.PostForm("use_password")
	fmt.Println("pass", password)

	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, _ := ks.NewAccount(password)
	data, err := config.RegisterMethod(client, contract, account.Address, name, email, password)

	fileInfoList,err :=ioutil.ReadDir("./keystore")
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("len(fileInfoList):",len(fileInfoList))
	for i:=range fileInfoList{
		if i ==len(fileInfoList)-1 {
			names =fileInfoList[i].Name()
			fmt.Println(names)
		}
	}

	keyjson, errs := ioutil.ReadFile("./keystore/" + names)
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
	fmt.Println("pri", privKey)
	key,err :=DumpPrivateKeyBase64(privKey)
	err1 :=config.EmailSend1(email,names,key)
	if err1 !=nil {
		fmt.Println(err1)
	}
	respOK(c, data)
	fmt.Println("data", data)
}

func DumpPrivateKeyBase64(privatekey *ecdsa.PrivateKey) (string, error) {
	keybytes,err := x509.MarshalECPrivateKey(privatekey)
	if err !=nil {
		fmt.Println(err)
	}
	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}