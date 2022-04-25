package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func topUp(c *gin.Context) {
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

	amount := c.PostForm("amount")
	fmt.Println("amount",amount)
	amountf, err := strconv.ParseFloat(amount, 64) //先转换为 float64

	if err != nil {

		log.Println("is not a number")

	}

	// 再通过sprintf格式化为*Int

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", amountf*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}
	fmt.Println("wei",valueWei)
	opts := config.Getopts()

	opts.Value = valueWei
	amountInt, err := strconv.Atoi(amount)
	amountInt64 := int64(amountInt)
	res, err := contract.TopUp(opts,big.NewInt(amountInt64),LoginUser)

	fmt.Println("topUp:", res)
	opts.GasLimit = 3000000
	opts.GasPrice, err = config.GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, "/profile")
}
