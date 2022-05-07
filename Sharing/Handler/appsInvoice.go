package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

type OrderDit struct {
	OrderOwner    common.Address //出借人
	OrderBorrower common.Address //租借者地址
	OrderId       int            //物品id
	Deal          int            //借用记录
	Back          int            //归还记录
}

//渲染con
func AppInvoiceStatic(c *gin.Context) {
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
	orderId := config.HaveOrderId(client)
	for i := 0; i < len(orderId); i++ {

	}
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)

	if err != nil {
		respError(c, err)
		return
	}
	OId := c.Params.ByName("OId")

	idInt, err := strconv.Atoi(OId)
	idInt64 := int64(idInt)
	OrderDet, _ := config.HaveOrderDit(client, big.NewInt(idInt64))
	goodData, goodData1, err2 := config.HaveIndex(client, OrderDet.Id)
	blockNum, timeStamp, dealHash, days, startTime, err := contract.GetDealRec(nil, OrderDet.Id, OrderDet.Deal)
	dealHashStr:=fmt.Sprintf("0x%x",dealHash)
	isBack, endTime, err := contract.GetBackRec(nil, OrderDet.Id, OrderDet.Back)
	if err2 != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/apps-invoice.html", gin.H{
		"id":        OrderDet.Id,
		"name":      goodData.Name,
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"owner":     OrderDet.OrderOwner,
		"borrower":  OrderDet.OrderBorrower,
		"goodImg":   goodData1.GoodImg,
		"rent":      goodData.Rent,
		"blockNum":  blockNum,
		"dealHash":  dealHashStr,
		"startTime": startTime,
		"endTime":   endTime,
		"back":      OrderDet.Back,
		"timeStamp": timeStamp,
		"isBack":    isBack,
		"days":      days,
		"OId":       OId,
	})
}
