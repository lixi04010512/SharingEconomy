package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

//租借者信息
type Renter struct {
	RenterAddr common.Address `json:renterAddr ` //租借者地址
	Since      int            `json:since`       //租借开始时间renter
	Over       int            `json:over`        //租借截止时间

}
type GoodsPort struct {
	Id        *big.Int       `json:"id"`
	borrower  Renter         `json:"renter"`    //租用者
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
type StickAll struct {
	Sticks         string `json:"sticks"`
	GoodsPortStick GoodsPort
}
type GoodsAllInter interface {
	goodsAll()
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
	var arr []GoodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var Owner common.Address
			var arr1 []GoodsPort
			if id[i] != big.NewInt(0) {
				goodData, goodData1, err := config.HaveIndex(client, id[i])
				//print("图片路径",goodImg)
				if err != nil {
					respError(c, err)
					return
				}
				arr1 = []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg}}

			}


			arr = append(arr, arr1...)

			//goodsPort1 := goodsPort{ Names: names, Species: species, Rent: rent, EthPledge: ethPledge}
			//fmt.Println(goodsPort{},addr)
		} else {
			userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
			userImg, err := contract.GetUserImg(nil, LoginUser)
			fmt.Println("res", userName)
			if err != nil {
				respError(c, err)
				return
			}
			c.HTML(http.StatusOK, "Static/index.html", gin.H{
				"goodsPor": arr,
				"userName": userName,
				"userImg":  userImg,
				"address":  people,
			})
		}
	}
}

//商品详情页展示
func shopPorduct(c *gin.Context) {
	a := c.Params.ByName("id")
	x, err := strconv.ParseInt(a, 10, 64)
	id := big.NewInt(x)
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

	//id := config.HaveId(client)
	//if err != nil {
	//	fmt.Println(err)
	//	respError(c, err)
	//	return
	//}
	var names string
	goodData, goodData1, err := config.HaveIndex(client, id)
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/shop-product.html", gin.H{
		"goodsPort": GoodsPort{
			Id:        goodData.Id,
			Names:     names,
			Rent:      goodData.Rent,
			EthPledge: goodData.EthPledge,
			GoodImg:   goodData1.GoodImg,
		},
		"userName": userName,
		"address":  people,
		"userImg":  userImg,
	})
}

//限制展示个数
var additionGoods int

func addGoodsCategory(int) int {
	additionGoods++
	if additionGoods >= 3 {
		additionGoods = 0
	}
	//fmt.Println("additionGoods",additionGoods)
	return additionGoods
}

//商品分类展示
func goodsCategory(c *gin.Context) {
	client, err := config.GetClient()
	contract, err := config.GetAddress(client)
	//获取id
	id := config.HaveId(client)
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
	if err != nil {
		respError(c, err)
		return
	}
	//定义一个结构体数组
	var stickArr []StickAll
	var arr []StickAll
	for j := 1; j < 9; j++ {
		if j < 8 {
			StickData, err := config.ShowStick(client, big.NewInt(int64(j)))
			if err != nil {
				respError(c, err)
				return
			}
			stickArr1 := []StickAll{StickAll{Sticks: StickData}}
			stickArr = append(stickArr, stickArr1...)
			for i := 1; i < len(id); i++ {
				goodData, goodData1, _ := config.HaveIndex(client, id[i])
				if goodData.Species == StickData {
					//arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg}}
					//arr = append(arr, arr1...)
					arr1 := []StickAll{StickAll{GoodsPortStick: GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg}}}
					stickArr = append(stickArr, arr1...)
					//fmt.Println("stickA融入33",stickArr)
					//goodsPort1 := goodsPort{ Names: names, Species: species, Rent: rent, EthPledge: ethPledge}
					//fmt.Println(goodsPort{},addr)
				}
			}
		} else {
			//fmt.Println("stickArr===", stickArr)
			fmt.Println("ashvdlasjvhlja", additionGoods)
			stickArr = append(stickArr, arr...)
			c.HTML(http.StatusOK, "Static/goods-category.html", gin.H{
				"goodsCategory": stickArr,
				"userName":      userName,
				"address":       people,
				"userImg":       userImg,
				"LimitAdd":      additionGoods,
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
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
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
	var arrUp []GoodsPort
	var arrDown []GoodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var addr common.Address
			goodData, goodData1, err := config.HaveIndex(client, id[i])

			if err != nil {
				respError(c, err)
				return
			}

			if people == goodData.Owner && goodData.Available == true {
				arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrUp = append(arrUp, arr1...)
			} else if people == goodData.Owner && goodData.Available == false {
				arr2 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrDown = append(arrDown, arr2...)
			}
		} else {

			c.HTML(http.StatusOK, "Static/cart.html", gin.H{
				"goodsUp":  arrUp,
				"goodDown": arrDown,
				"userName": userName,
				"address":  people,
				"userImg":  userImg,
			})
		}
	}
}

//我借出页
func MyOrder(c *gin.Context) {
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
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
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
	var arrUp []GoodsPort
	var arrDown []GoodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var addr common.Address
			goodData, goodData1, err := config.HaveIndex(client, id[i])
			if err != nil {
				respError(c, err)
				return
			}

			if people == goodData.Owner && goodData.IsBorrow == true {
				arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrUp = append(arrUp, arr1...)
			} else if people == goodData.Owner && goodData.IsBorrow == false {
				arr2 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg}}
				arrDown = append(arrDown, arr2...)
			}
		} else {

			c.HTML(http.StatusOK, "Static/order.html", gin.H{
				"goodsUp":  arrUp,
				"goodDown": arrDown,
				"userName": userName,
				"address":  people,
				"userImg":  userImg,
			})
		}
	}
}

//分类详情页面
func CategoryDetails(c *gin.Context) {
	Spe := c.Params
	Speci := Spe.ByName("Species")
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
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
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
	var arrDetails []GoodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var addr common.Address
			goodData, goodData1, err := config.HaveIndex(client, id[i])
			if err != nil {
				respError(c, err)
				return
			}

			if goodData.Species == Speci {
				arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, GoodImg: goodData1.GoodImg}}
				arrDetails = append(arrDetails, arr1...)
			}
			print(arrDetails)
		} else {

			c.HTML(http.StatusOK, "Static/category-details.html", gin.H{
				"categoryDetail": arrDetails,
				"userName":       userName,
				"address":        people,
				"userImg":        userImg,
			})
		}
	}
}

//我租入的物品页面
func Myshop(c *gin.Context) {
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
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
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
	var myshoparr []GoodsPort
	for i := 0; i < len(id)+1; i++ {
		if i < len(id) {
			//var addr common.Address
			goodData, goodData1, err := config.HaveIndex(client, id[i])
			if err != nil {
				respError(c, err)
				return
			}
			if goodData.Borrowers.Borrower == people && goodData.IsBorrow == true {
				arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, GoodImg: goodData1.GoodImg, Species: goodData.Species}}
				myshoparr = append(myshoparr, arr1...)
			}

		} else {

			c.HTML(http.StatusOK, "Static/shop.html", gin.H{
				"shopDet":  myshoparr,
				"userName": userName,
				"address":  people,
				"userImg":  userImg,
			})
		}
	}
}
