package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

//发送归还物品消息
func ReturnGoods(c *gin.Context) {
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
	id := c.PostForm("backId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	res,err := config.DoGoodsReturnMethod(client,contract,big.NewInt(idInt64),privKey)

	fmt.Println("backGoods",res)
	c.Redirect(http.StatusFound, "/cart")
}

//同意归还
//func AgreeBackGoods(c *gin.Context) {
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
//	res,err := config.AgreeBackMethod(client,contract,big.NewInt(idInt64),big.NewInt(borrowDaysInt64),privKey)
//
//	fmt.Println("sendBorrow",res)
//	c.Redirect(http.StatusFound, "/cart")
//}

//归还物品
func Back(c *gin.Context) {
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
	backs := c.PostForm("backs")
	backsInt, err := strconv.Atoi(backs)
	backsInt64 := int64(backsInt)
	over :=c.PostForm("over")
	res,err := config.BackGoodsMethod(client,contract,big.NewInt(idInt64),big.NewInt(backsInt64),over,privKey)

	fmt.Println("Back",res)
	c.Redirect(http.StatusFound, "/cart")
}
