package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

type goodsPort struct {
	Addr      string   `json:"addr"`      //地址
	Names     string   `json:"names"`     //名称
	Species   string   `json:"species"`   //种类
	Rent      string   `json:"rent"`      //价格
	EthPledge string   `json:"ethPledge"` //押金
	GoodImg   []string `json:"GoodImg"`   //图片路径
	count     int      `json:"count"`     //借用次数
	deal      int      `json:"deal"`      //交易记录
	backs     int      `json:"backs"`     //归还记录
	available bool     `json:"available"` //是否上架
	isBorrow  bool     `json:"isBorrow"`  //是否借出

}

//首页-动态展示
func addIndex(c *gin.Context) {

	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	//获取id
	id := config.HaveId(client)
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}

	//定义一个结构体数组
	var arr []goodsPort
	var names, species, rent, ethPledge string
	var goodImg []string
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var Owner common.Address
			var Rent *big.Int
			var EthPledge *big.Int
			_, names, species, Rent, EthPledge, goodImg, err = config.HaveIndex(client, id[i])
			//addr = Owner.Hex()
			rent = Rent.String()
			ethPledge = EthPledge.String()
			//print("图片路径",goodImg)
			if err != nil {
				respError(c, err)
				return
			}
			arr1 := []goodsPort{goodsPort{Names: names, Species: species, Rent: rent, EthPledge: ethPledge, GoodImg: goodImg}}
			arr = append(arr, arr1...)

			//goodsPort1 := goodsPort{ Names: names, Species: species, Rent: rent, EthPledge: ethPledge}
			//fmt.Println(goodsPort{},addr)
		} else {
			userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
			fmt.Println("res", userName)
			if err != nil {
				respError(c, err)
				return
			}
			c.HTML(http.StatusOK, "Static/index.html", gin.H{
				"goodsPor": arr,
				"userName": userName,
				"address":  people,
			})
		}
	}
}

//商品详情页展示
func shopPorduct(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	id := config.HaveId(client)
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}

	var names string
	var goodImg []string
	var Rent, EthPledge *big.Int
	_, names, _, Rent, EthPledge, goodImg, err = config.HaveIndex(client, id[0])
	rent := Rent.String()
	ethPledge := EthPledge.String()

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	//arr1 :=[]goodsPort{ goodsPort{Names: names,Rent: rent,EthPledge: ethPledge}}
	//arr=append(arr, arr1...)

	c.HTML(http.StatusOK, "Static/shop-product.html", gin.H{
		"goodsPort": goodsPort{
			Names:     names,
			Rent:      rent,
			EthPledge: ethPledge,
			GoodImg:   goodImg,
		},
		"userName": userName,
		"address":  people,
	})
}

//商品分类展示
func goodsCategory(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	//获取id
	id := config.HaveId(client)
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//定义一个结构体数组
	var arr []goodsPort
	var names, species, rent, ethPledge string
	var goodImg []string
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var Owner common.Address
			var Rent *big.Int
			var EthPledge *big.Int
			_, names, species, Rent, EthPledge, goodImg, err = config.HaveIndex(client, id[i])
			//addr = Owner.Hex()
			rent = Rent.String()
			ethPledge = EthPledge.String()
			//print("图片路径",goodImg)
			if err != nil {
				respError(c, err)
				return
			}
			arr1 := []goodsPort{goodsPort{Names: names, Species: species, Rent: rent, EthPledge: ethPledge, GoodImg: goodImg}}
			arr = append(arr, arr1...)

			//goodsPort1 := goodsPort{ Names: names, Species: species, Rent: rent, EthPledge: ethPledge}
			//fmt.Println(goodsPort{},addr)
		} else {
			userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
			fmt.Println("res", userName)
			if err != nil {
				respError(c, err)
				return
			}
			c.HTML(http.StatusOK, "Static/goods-category.html", gin.H{
				"goodsCategory": arr,
				"userName":      userName,
				"address":       people,
			})
		}
	}
}
