package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

//上架物品
func addGoods(c *gin.Context) {
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
	form, err := c.MultipartForm()
	name := c.PostForm("goodsName")
	owner := LoginUser
	species := c.PostForm("species")
	rent := c.PostForm("rent")
	//string转*big.int
	rentInt, err := strconv.Atoi(rent)
	rentInt64 := int64(rentInt)

	ethPledge := c.PostForm("textDeposit")
	ethPledgeInt, err := strconv.Atoi(ethPledge)
	ethPledgeInt64 := int64(ethPledgeInt)
	goodSign := c.PostForm("addnote")


	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files := form.File["myGoodsImg"]
	fmt.Println("files",files)

	//存储所有图片路径
	var goodsImgs []string

	fildDir := "./Static/images/"

	// 遍历所有图片
	for _, file := range files {
		fmt.Println("fileok")
		// 逐个存
		fileName :=file.Filename
		filepath := fmt.Sprintf("%s%s", fildDir, fileName)
		fileH :=fmt.Sprintf("share/images/%s",fileName)
		//fmt.Println("path",filepath)
		goodsImgs = append(goodsImgs, fileH)
		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	//goodsImgs := c.PostForm("goodsImgs")
	fmt.Println("pass", name)
	fmt.Println("goodsImg",goodsImgs)

	data, err := config.AddGoodsMethod(client, contract, owner, name, species, big.NewInt(rentInt64), big.NewInt(ethPledgeInt64), goodsImgs,goodSign,privKey)
	fmt.Println("addGood",data)
	c.Redirect(http.StatusFound, "/cart")
}

