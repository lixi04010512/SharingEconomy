package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

func delGoods(c *gin.Context) {
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

	id := c.PostForm("delId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	goodsData,_, err := config.HaveIndex(client, big.NewInt(idInt64))
	data,err := config.OutGoodsMethod(client,contract,goodsData.Owner,big.NewInt(idInt64),privKey)
	fmt.Println("data",data)
}
