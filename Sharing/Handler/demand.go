package handler

import (
	config "Sharing/Config"
	x "Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//增加需求
func demandAdd(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println()
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	needs := x.DemandDB{}
	demandKinds := c.PostForm("demandKinds")
	demandNames := c.PostForm("demandName")

	needs.DemandKinds = demandKinds
	demandAddr := people.Hex()
	//fmt.Println("我的需求拦截器", demandAddr)
	err = x.InsertDemand(demandKinds, demandAddr, demandNames)
	if demandAddr == "0x0000000000000000000000000000000000000000" {
		respError(c, "请先登录")
	} else {
		respOK(c, "ok")
	}

}

//查
func ChatStatic(c *gin.Context) {
	//listNeeds, err := x.ListNeedAll()
	//fmt.Println(listNeeds)
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println()
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/chat.html", gin.H{
		"userName": userName,
		"address":  people,
		"userImg":  userImg,
	})
}
