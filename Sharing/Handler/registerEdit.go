package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"
)

//邮箱验证
func sendEmail(c *gin.Context) {
	email := c.PostForm("use_email")
	code := config.EmailSend(email)
	respOK(c, code)
}

//注册
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
	respOK(c, data)
	fmt.Println("data", data)
}
