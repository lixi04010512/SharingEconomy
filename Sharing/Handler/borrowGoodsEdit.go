package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

//发送借用物品消息
func BorrowGoods(c *gin.Context) {
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
	id := c.PostForm("borrowId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	borrowDays := c.PostForm("borrowDays")
	borrowDaysInt, err := strconv.Atoi(borrowDays)
	borrowDaysInt64 := int64(borrowDaysInt)
	res,err := config.BorrowGoodsMethod(client,contract,big.NewInt(idInt64),big.NewInt(borrowDaysInt64),privKey)

	fmt.Println("sendBorrow",res)
	c.Redirect(http.StatusFound, "/cart")
}

//发送借用物品消息
//func AgreeBorrowGoods(c *gin.Context) {
//	//初始化client
//	client, err := config.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	//初始化合约地址
//	contract, err := config.GetAddress(client)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//	id := c.PostForm("borrowId")
//	idInt, err := strconv.Atoi(id)
//	idInt64 := int64(idInt)
//	borrowDays := c.PostForm("borrowDays")
//	borrowDaysInt, err := strconv.Atoi(borrowDays)
//	borrowDaysInt64 := int64(borrowDaysInt)
//	res,err := config.AgreeMethod(client,contract,big.NewInt(idInt64),big.NewInt(borrowDaysInt64),privKey)
//
//	fmt.Println("sendBorrow",res)
//	c.Redirect(http.StatusFound, "/cart")
//}


