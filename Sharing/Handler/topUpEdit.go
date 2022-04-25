package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

func topUp(c *gin.Context) {
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

	amount := c.PostForm("amount")
	fmt.Println("amount",amount)
	amountInt, err := strconv.Atoi(amount)
	amountInt64 := int64(amountInt)
	data,err := config.TopUpMethod(client,contract,big.NewInt(amountInt64),LoginUser)
	fmt.Println("data",data)
	c.Redirect(http.StatusFound, "/profile")
}
