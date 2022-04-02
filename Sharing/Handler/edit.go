package handler

import (
	"Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)
//用户地址
var loginUser common.Address

func addSticker(c *gin.Context) {
	//初始化client
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
	//初始化opts
	opts := config.Getopts()

	//初始化gas
	gasPrice, err := config.GetgasPrice(client)
	opts.GasLimit = 3000000
	opts.GasPrice = gasPrice

	stick := c.PostForm("stick")
	val, err := contract.AddSticker(opts, stick)
	fmt.Println(val)
	respOK(c, val)

}

func isStickExist(c *gin.Context) {
	//初始化client
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

	stick := c.PostForm("stick")
	val, err := contract.IsStickExist(nil, stick)
	fmt.Println(val)
	respOK(c, val)
}

//邮箱验证
func sendEmail(c *gin.Context)  {
	email :=c.PostForm("use_email")
	code:=config.EmailSend(email)
	respOK(c,code)
}

//注册
func register(c *gin.Context){
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

	name := c.PostForm("use_name")
	email :=c.PostForm("use_email")
	password := c.PostForm("use_password")
	fmt.Println("pass",password)

	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account, _ := ks.NewAccount(password)
	data,err:=config.RegisterMethod(client,contract,account.Address,name,email,password)
	respOK(c, data)
	fmt.Println("data",data)
}
//登录
func login(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		fmt.Println(err)
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	_,f, err := c.Request.FormFile("log_addr")
	if err != nil{
		fmt.Println(err)
			respError(c,err)
			return
		}
	password := c.PostForm("log_passwd")

		keyjson,errs:=ioutil.ReadFile("./keystore/"+f.Filename)
		fmt.Println("keyjson:",string(keyjson))
		if errs != nil{
			respError(c,err)
			return
		}
		unlockedKey, errors := keystore.DecryptKey(keyjson, password)
		fmt.Println("unlock:",unlockedKey)
		if errors != nil{
			respError(c,err)
			fmt.Println(errors)
			return
		}
		privKey :=unlockedKey.PrivateKey
		comAddr := unlockedKey.Address
		loginUser =comAddr
		data,err:=config.LoginMethod(client,contract,comAddr,privKey)
		addr :=comAddr.Hex()

				c.Redirect(http.StatusFound, "/index")
				//respOK(c,data)
				fmt.Println("addr", addr)
				fmt.Println("data", data)
	}

//注销
func logout(c *gin.Context)  {
	//初始化client
	client,err := config.GetClient()
	if err != nil{
		fmt.Println(err)
		respError(c,err)
		return
	}
	//初始化合约地址
	contract ,err:= config.GetAddress(client)
	if err != nil{
		respError(c,err)
		return
	}

	data,err:=contract.Logout(nil,loginUser)
	loginUser=common.Address{}
	respOK(c,data)
}

//上架物品
func addGoods(c *gin.Context)  {
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

	name := c.PostForm("name")
	owner:=loginUser
	species :=c.PostForm("species")
	rent := c.PostForm("rent")
	//string转*big.int
	rentInt,err:=strconv.Atoi(rent)
	rentInt64 := int64(rentInt)

	ethPledge := c.PostForm("ethPledge")
	ethPledgeInt,err:=strconv.Atoi(ethPledge)
	ethPledgeInt64 := int64(ethPledgeInt)

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files := form.File["files"]

	//存储所有图片路径
	var goodsImgs []string

	fildDir:="./Static/images"


	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		fileName:=file.Filename
		filepath:=fmt.Sprintf("%s%s",fildDir,fileName)
		goodsImgs=append(goodsImgs,filepath)
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	//goodsImgs := c.PostForm("goodsImgs")
	fmt.Println("pass",name)
	data,err := config.AddGoodsMethod(client,contract,owner,name,species,big.NewInt(rentInt64),big.NewInt(ethPledgeInt64),goodsImgs)
	respOK(c,data)
}

//修改用户信息
//func updateUser(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		respError(c,err)
//		fmt.Println(err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//}