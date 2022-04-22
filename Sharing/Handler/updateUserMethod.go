package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//修改用户信息
func updateUser(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		respError(c,err)
		fmt.Println(err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	name := c.PostForm("up_name")
	email := c.PostForm("up_email")
	sign := c.PostForm("up_sign")
	fmt.Println("sign", sign)

	data, err := config.UpdateUserMethod(client, contract, LoginUser, name, email, sign,privKey)
	c.Redirect(http.StatusFound, "/profile")
	fmt.Println("data", data)

}
