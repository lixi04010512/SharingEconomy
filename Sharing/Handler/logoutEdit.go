package handler

import (
	config "Sharing/Config"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logoutUser common.Address
var logoutPriKey *ecdsa.PrivateKey
//注销
func logout(c *gin.Context) {
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

	data, err := config.LogoutMethod(client,contract, LoginUser,privKey)
	fmt.Println("注销",data)
	LoginUser = logoutUser
	privKey =logoutPriKey
	c.Redirect(http.StatusFound, "/login")
}
