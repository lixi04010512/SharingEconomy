package handler

import (
	config "Sharing/Config"
	"Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListNeedAll(c *gin.Context) {
	x, err := db.ListNeedAll()
	if err != nil {
		respError(c, err)
		return
	}
	//fmt.Println("x----x",x)
	respOK(c, x)
}

func MYList(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		respError(c, err)
		return
	}
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}
	_, people, _, _, _, _, _, err := config.GetUserMethod(contract, LoginUser)

	if err != nil {
		respError(c, err)
		return
	}
	x, err := db.MyListNeeds(people)
	fmt.Println("demandID", x)
	respOK(c, x)
}
func delNeeds(c *gin.Context) {
	a := c.Params.ByName("id")
	id, err := strconv.ParseInt(a, 10, 64)
	x, err := db.DeleteNeeds(id)
	fmt.Println(x)
	if err != nil {
		respError(c, err)
		return
	}

	c.Redirect(http.StatusFound, "../myneed")
}
