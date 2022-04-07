package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

type goodsPort struct {
	Id        *big.Int       `json:"id"`
	Addr      common.Address `json:"addr"`      //地址
	Names     string         `json:"names"`     //名称
	Species   string         `json:"species"`   //种类
	Rent      *big.Int       `json:"rent"`      //价格
	EthPledge *big.Int       `json:"ethPledge"` //押金
	GoodImg   []string       `json:"GoodImg"`   //图片路径
	Count     int            `json:"count"`     //借用次数
	Deal      int            `json:"deal"`      //交易记录
	Backs     int            `json:"backs"`     //归还记录
	Available bool           `json:"available"` //是否上架
	IsBorrow  bool           `json:"isBorrow"`  //是否借出

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
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var Owner common.Address

			goodData, goodData1, err := config.HaveIndex(client, id[i])

			//print("图片路径",goodImg)
			if err != nil {
				respError(c, err)
				return
			}
			arr1 := []goodsPort{goodsPort{Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg}}
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
	goodData, goodData1, err := config.HaveIndex(client, id[0])
	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/shop-product.html", gin.H{
		"goodsPort": goodsPort{
			Id:        goodData.Id,
			Names:     names,
			Rent:      goodData.Rent,
			EthPledge: goodData.EthPledge,
			GoodImg:   goodData1.GoodImg,
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

	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			goodData, goodData1, err := config.HaveIndex(client, id[i])
			if err != nil {
				respError(c, err)
				return
			}
			arr1 := []goodsPort{goodsPort{Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg}}
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

//我的交易-上下架
func CartGood(c *gin.Context) {
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
	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
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
	var arrUp []goodsPort
	var arrDown []goodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var addr common.Address
			goodData, goodData1, err := config.HaveIndex(client, id[i])
			if err != nil {
				respError(c, err)
				return
			}

			if people == goodData.Owner && goodData.Available == true {
				arr1 := []goodsPort{goodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrUp = append(arrUp, arr1...)
			} else if people == goodData.Owner && goodData.Available == false {
				arr2 := []goodsPort{goodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrDown = append(arrDown, arr2...)
			}
		} else {

			c.HTML(http.StatusOK, "Static/cart.html", gin.H{
				"goodsUp":  arrUp,
				"goodDown": arrDown,
				"userName": userName,
				"address":  people,
			})
		}
	}
}
