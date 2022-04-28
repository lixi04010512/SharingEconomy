package handler

import (
	config "Sharing/Config"
	"Sharing/db"
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
	id := c.PostForm("id")
	fmt.Println("id",id)
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	borrowDays := c.PostForm("borrowDays")
	fmt.Println("bo",borrowDays)
	borrowDaysInt, err := strconv.Atoi(borrowDays)
	borrowDaysInt64 := int64(borrowDaysInt)
	res, deal, _, err := config.BorrowGoodsMethod(client, contract, big.NewInt(idInt64), big.NewInt(borrowDaysInt64), privKey)
	message := fmt.Sprintf("你好，我是账号为%s的用户，我想借用%s天你的商品，我的借用号是%s，", LoginUser, borrowDays, deal)
	goodsData,_, err := config.HaveIndex(client, big.NewInt(idInt64))
	goods_owner := goodsData.Owner.String()
	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img,err :=contract.GetUserImg(nil, LoginUser)
	userName, _, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	name_owner, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Owner)
	fmt.Println("mess", message)
	addr := LoginUser.Hex()
	err = db.SendBorrow(addr, goods_owner, userName, name_owner, message, userImg, img)
	if err != nil {
		respError(c, err)
	}

	fmt.Println("sendBorrow", res)
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
//
//
//	fmt.Println("sendBorrow",res)
//	c.Redirect(http.StatusFound, "/cart")
//}
//
////借用物品
//func Borrow(c *gin.Context) {
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
//	deal := c.PostForm("deal")
//	dealInt, err := strconv.Atoi(deal)
//	dealInt64 := int64(dealInt)
//	res,err := config.AgreeMethod(client,contract,big.NewInt(idInt64),big.NewInt(borrowDaysInt64),privKey)
//	opts := config.GetMsgOpts(privateKey)
//	amountf, err := strconv.ParseFloat(amount, 64) //先转换为 float64
//
//	if err != nil {
//
//		log.Println("is not a number")
//
//	}
//
//	// 再通过sprintf格式化为*Int
//
//	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", amountf*1000000000000000000), 10)
//
//	if !isOk {
//		log.Println("float to bigInt failed!")
//	}
//	opts.Value = valueWei
//	opts.GasLimit = 3000000
//	opts.GasPrice, err = config.GetgasPrice(client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, err := contract.Borrow(opts, big.NewInt(idInt64), big.NewInt(dealInt64))
//	fmt.Println("borrow:", res)
//	c.Redirect(http.StatusFound, "/cart")
//}
