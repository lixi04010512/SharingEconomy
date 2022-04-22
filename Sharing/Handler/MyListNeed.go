package handler

import (
	config "Sharing/Config"
	"Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
	fmt.Println("people", people, x)
	respOK(c, x)
}
