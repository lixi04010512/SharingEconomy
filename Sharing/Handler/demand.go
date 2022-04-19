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
	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	needs := x.Demand{}
	demandKinds := c.PostForm("demandKind")
	demandNames := c.PostForm("demandName")
	fmt.Println(demandKinds, "====", demandNames)
	needs.DemandKind = demandKinds
	demandAddr := people.Hex()
	err = x.InsertDemand(demandKinds, demandAddr, demandNames)
}

//查
func ChatStatic(c *gin.Context) {
	listNeeds, err := x.ListNeeds()
	fmt.Println(listNeeds)
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
	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/chat.html", gin.H{
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"listNeeds": listNeeds})
}
