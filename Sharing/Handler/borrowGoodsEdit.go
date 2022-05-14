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
	id1 := c.PostForm("id")
	fmt.Println("id", id1)
	idInt, err := strconv.Atoi(id1)
	idInt64 := int64(idInt)
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))
	rent := goodsData.Rent
	rentStr := rent.String()              //转成string
	rentInt, err := strconv.Atoi(rentStr) //string转int
	borrowDays := c.PostForm("borrowDays")
	fmt.Println("bo", borrowDays)
	borrowDaysInt, err := strconv.Atoi(borrowDays)
	borrowDaysInt64 := int64(borrowDaysInt)
	pledge := goodsData.EthPledge
	pledgeStr := pledge.String()              //转成string
	pledgeInt, err := strconv.Atoi(pledgeStr) //string转int
	transfEth := rentInt*borrowDaysInt + pledgeInt
	transfEths := strconv.Itoa(transfEth)
	amountf, err := strconv.ParseFloat(transfEths, 64) //先转换为 float64
	//
	if err != nil {

		log.Println("is not a number")

	}

	// 再通过sprintf格式化为*Int

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", amountf*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}
	res, deal, _, err := config.BorrowGoodsMethod(client, contract, big.NewInt(idInt64), big.NewInt(borrowDaysInt64), valueWei, privKey)
	message := fmt.Sprintf("你好，我是账号为%s的用户，我想借用你的%s号商品%s天，本次交易id是%s，", LoginUser, id1, borrowDays, deal)

	goods_owner := goodsData.Owner.String()
	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, LoginUser)
	userName, _, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	name_owner, _, _, _, _, err := config.GetUserMethod(contract, goodsData.Owner)
	fmt.Println("mess", message)
	addr := LoginUser.Hex()
	userImg1 := "share/" + userImg
	img1 := "share/" + img
	dealString := deal.String()
	dealInt, err := strconv.Atoi(dealString)
	err = db.SendBorrow(addr, goods_owner, userName, name_owner, message, img1, userImg1, idInt, dealInt)
	if err != nil {
		respError(c, err)
	}

	fmt.Println("sendBorrow", res)
}

//不同意借出
func DisagreeBorrowGoods(c *gin.Context) {
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
	id := c.PostForm("disagreeId")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	deal := c.PostForm("dealId")
	dealInt, err := strconv.Atoi(deal)
	dealInt64 := int64(dealInt)
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))
	rent := goodsData.Rent
	rentStr := rent.String()              //转成string
	rentInt, err := strconv.Atoi(rentStr) //string转int
	_, _, _, borrowDays, _, err := contract.GetDealRec(nil, big.NewInt(idInt64), big.NewInt(dealInt64))
	borrowDaysStr := borrowDays.String()              //转成string
	borrowDaysInt, err := strconv.Atoi(borrowDaysStr) //string转int
	pledge := goodsData.EthPledge
	pledgeStr := pledge.String()              //转成string
	pledgeInt, err := strconv.Atoi(pledgeStr) //string转int

	transfEth := rentInt*borrowDaysInt + pledgeInt
	transfEths := strconv.Itoa(transfEth)
	amountf, err := strconv.ParseFloat(transfEths, 64) //先转换为 float64
	if err != nil {

		log.Println("is not a number")

	}

	// 再通过sprintf格式化为*Int

	value, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", amountf*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}

	res, err := config.DisagreeMethod(client, contract, big.NewInt(idInt64), big.NewInt(dealInt64), value, privKey)

	//发送留言消息
	message := c.PostForm("message")
	fmt.Println("mess", message)
	fmt.Println("sendDisBorrow", res)

	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, goodsData.Borrowers.Borrower)
	userImg1 := "share/" + userImg
	img1 := "share/" + img
	fmt.Println("disagress:", goodsData.Borrowers.Borrower.String())
	err1 := db.DisagreeBorrow(goodsData.Owner.String(), goodsData.Borrowers.Borrower.String(), message, userImg1, img1)
	if err1 != nil {
		fmt.Println("146err1:", err1)
	}
}

//同意借用物品
func AgreeBorrow(c *gin.Context) {
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
	deal := c.PostForm("dealId")
	dealInt, err := strconv.Atoi(deal)
	dealInt64 := int64(dealInt)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	goodsData, _, err := config.HaveIndex(client, big.NewInt(idInt64))

	blockNum, _, _, _, _, err := contract.GetDealRec(nil, big.NewInt(idInt64), big.NewInt(dealInt64))

	res, err := config.AgreeBorrowMethod(client, contract, big.NewInt(idInt64), big.NewInt(dealInt64), timeStr, privKey)

	dealHash, err := config.HashMethod(client, contract, big.NewInt(idInt64), big.NewInt(dealInt64), blockNum, privKey)

	//发送留言消息
	message := c.PostForm("message")
	fmt.Println("mess", message)
	fmt.Println("borrow:", res)
	fmt.Println("hash", dealHash)

	userImg, err := contract.GetUserImg(nil, goodsData.Owner)
	img, err := contract.GetUserImg(nil, goodsData.Borrowers.Borrower)
	fmt.Println("goodsData.Owner.String():", goodsData.Owner.String())
	fmt.Println("goodsData.Borrowers.Borrower.String()", goodsData.Borrowers.Borrower.String())
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
