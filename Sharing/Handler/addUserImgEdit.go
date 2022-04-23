package handler

import (
	"Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//添加头像
func AddUserImg(c *gin.Context)  {
	fmt.Println("ok")
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

	img,err :=config.UploadUserImg(c)
	if err != nil{
		respError(c,err)
		return
	}
	fmt.Println("img",img)
	data,err :=config.AddUserImgMethod(client,contract, LoginUser,img,privKey)
	fmt.Println("userImg",data)
	c.Redirect(http.StatusFound, "/profile")

}
