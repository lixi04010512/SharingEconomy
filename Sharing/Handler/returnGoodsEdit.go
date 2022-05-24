package handler

import (
	config "Sharing/Config"
	"Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strconv"
	"time"
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
	id := c.PostForm("id")
	fmt.Println("id", id)
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))

	res, back, err := config.DoGoodsReturnMethod(client, contract, big.NewInt(idInt64), privKey)
	message := fmt.Sprintf("你好，我是账号为%s的用户，我要归还你的%s号商品，本次归还id是%s，", LoginUser, id, back)

	goods_owner := goodsData.Owner.String()
	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, LoginUser)
	userName, _, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	name_owner, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Owner)
	fmt.Println("mess", message)
	addr := LoginUser.Hex()
	userImg1 := "share/" + userImg
	img1 := "share/" + img
	err = db.SendBorrow(addr, goods_owner, userName, name_owner, message, userImg1, img1)
	if err != nil {
		respError(c, err)
	}

	fmt.Println("sendBorrow", res)
}

//同意归还
func AgreeBackGoods(c *gin.Context) {
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
	//页面获取id
	id := c.PostForm("agreeId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	back := c.PostForm("backId")
	backInt, err := strconv.Atoi(back)
	backInt64 := int64(backInt)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))
	pledge := goodsData.EthPledge
	pledgeStr := pledge.String() //转成string

	//isBack,err:=contract.GetBackRec(nil,big.NewInt(idInt64),big.NewInt(backInt64))

	//获取转账金额
	amountf, err := strconv.ParseFloat(pledgeStr, 64) //先转换为 float64
	if err != nil {

		log.Println("is not a number")

	}

	// 再通过sprintf格式化为*Int

	value, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", amountf*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}
	res, err := config.AgreeBackMethod(client, contract, big.NewInt(idInt64), big.NewInt(backInt64), value, timeStr, privKey)

	//发送留言消息
	message := c.PostForm("message")
	fmt.Println("mess", message)
	fmt.Println("borrow:", res)
	fmt.Println("value", value)

	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, goodsData.Borrowers.Borrower)
	userName, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Borrowers.Borrower)
	name_owner, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Owner)
	fmt.Println("goodsData.Owner.String():", goodsData.Owner.String())
	fmt.Println("goodsData.Borrowers.Borrower.String()", goodsData.Borrowers.Borrower.String())
	fmt.Println("name_owner", name_owner)
	fmt.Println("userName", userName)
	fmt.Println("message", message)
	fmt.Println("userImg", userImg)
	fmt.Println("img", img)
	userImg1 := "share/" + userImg
	img1 := "share/" + img
	err1 := db.AgreeBorrow(goodsData.Owner.String(), goodsData.Borrowers.Borrower.String(), message, userImg1, img1)
	if err1 != nil {
		fmt.Println("202err:", err1)
	}
}

//归还物品
func DisagreeBackGoods(c *gin.Context) {
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
	//页面获取id
	id := c.PostForm("disagreeId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	back := c.PostForm("backId")
	backInt, err := strconv.Atoi(back)
	backInt64 := int64(backInt)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))
	res, err := config.DisagreeBackMethod(client, contract, big.NewInt(idInt64), big.NewInt(backInt64), timeStr, privKey)
	//_,_,err:=contract.GetBackRec(nil,big.NewInt(idInt64),big.NewInt(backInt64))

	//发送留言消息
	message := c.PostForm("message")
	fmt.Println("mess", message)
	fmt.Println("borrow:", res)

	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, goodsData.Borrowers.Borrower)
	userName, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Borrowers.Borrower)
	name_owner, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Owner)
	fmt.Println("goodsData.Owner.String():", goodsData.Owner.String())
	fmt.Println("goodsData.Borrowers.Borrower.String()", goodsData.Borrowers.Borrower.String())
	fmt.Println("name_owner", name_owner)
	fmt.Println("userName", userName)
	fmt.Println("message", message)
	fmt.Println("userImg", userImg)
	fmt.Println("img", img)
	userImg1 := "share/" + userImg
	img1 := "share/" + img
	err1 := db.DisagreeBorrow(goodsData.Owner.String(), goodsData.Borrowers.Borrower.String(), message, userImg1, img1)
	if err1 != nil {
		fmt.Println("202err:", err1)
	}
}
