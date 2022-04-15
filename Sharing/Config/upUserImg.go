package config

import (
	"Sharing/Agreement"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strings"
)

//图片上传
func UploadUserImg(c *gin.Context) (string,error){
	f, err := c.FormFile("imgname")
	fileName := f.Filename
	fildDir := "./Static/images/"

	filePath := fmt.Sprintf("%s%s", fildDir, fileName)
	fileH :=fmt.Sprintf("images/%s",fileName)
	//file, handler, err := c.FormFile("imgname")
	fmt.Println(err)
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,jpeg文件",
			})
			return "",err
		}

		c.SaveUploadedFile(f, filePath)
		fmt.Println("filename", fileName)

	}
	return fileH,err
}

//封装头像上传方法
func AddUserImgMethod(client *ethclient.Client, contract *Agreement.User, people common.Address, headImg string,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.AddUserImg(opts, people,headImg)
	fmt.Println("addUserImg:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
