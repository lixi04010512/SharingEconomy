package handler

import (
	config "Sharing/Config"
	"Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func orderReceiving(c *gin.Context) {
	a := c.Params.ByName("id")
	id, err := strconv.ParseInt(a, 10, 64)
	client, err := config.GetClient()
	if err != nil {
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

	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	x, err := db.IdListNeeds(id)
	fmt.Println("需求", x)
	c.HTML(http.StatusOK, "Static/order-receiving.html", gin.H{
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"ListNeeds": x,
	})
}
