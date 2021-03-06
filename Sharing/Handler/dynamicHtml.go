package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//租借者信息
type Renter struct {
	RenterAddr common.Address `json:renterAddr ` //租借者地址
	Since      int            `json:since`       //租借开始时间renter
	Over       int            `json:over`        //租借截止时间

}
type GoodsPort struct {
	Days      int
	AllPrice  int
	IsBack    bool
	EndTime   string
	BackTime  string
	OId       *big.Int
	Id        *big.Int       `json:"id"`
	borrower  Renter         `json:"renter"`    //租用者
	Addr      common.Address `json:"addr"`      //地址
	Names     string         `json:"names"`     //名称
	Species   string         `json:"species"`   //种类
	Rent      *big.Int       `json:"rent"`      //价格
	EthPledge *big.Int       `json:"ethPledge"` //押金
	GoodImg   []string       `json:"GoodImg"`   //图片路径
	GoodSign  string         `json:"GoodSign"`  //商品描述
	Count     int            `json:"count"`     //借用次数
	Deal      int            `json:"deal"`      //交易记录
	Backs     int            `json:"backs"`     //归还记录
	Available bool           `json:"available"` //是否上架
	IsBorrow  bool           `json:"isBorrow"`  //是否借出

}
type StickAll struct {
	Sticks         string `json:"sticks"`
	SticksImg      string `json:"sticksImg"'`
	GoodsPortStick GoodsPort
}
type BaseController struct {
}

func (b BaseController) dynamicHtmlFun(c *gin.Context, idx int, id []int64, Speci string, webPage string) (string, common.Address, string, []GoodsPort, []GoodsPort, []GoodsPort, []GoodsPort) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
	}
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
	}
	var arrIndex []GoodsPort
	var arrUp []GoodsPort
	var arrDown []GoodsPort
	var arrDetails []GoodsPort
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
	for i := 0; i < idx; i++ {
		//var Owner common.Address
		//xint:=int64(rand.Intn(len(id)))
		//xint := int64(gRand.Intn(len(id)))
		goodData, goodData1, err3 := config.HaveIndex(client, big.NewInt(id[i]))
		//print("图片路径",goodImg)
		if err3 != nil {
			fmt.Println(err3)
		}
		//首页展示
		if webPage == "index" {
			if goodData.Name != "" && goodData.Available == true && goodData.IsBorrow == false {
				arrIndex1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
				arrIndex = append(arrIndex, arrIndex1...)
			}
		} else if webPage == "cart" || webPage == "order" {
			//借出页展示
			if people == goodData.Owner && goodData.Available == true {
				arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
				arrUp = append(arrUp, arr1...)
			} else if people == goodData.Owner && goodData.Available == false {
				arr2 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
				arrDown = append(arrDown, arr2...)
			}
		} else if webPage == "categoryDetails" {
			//分类详情页
			if goodData.Species == Speci {
				arrDetails1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, GoodImg: goodData1.GoodImg}}
				arrDetails = append(arrDetails, arrDetails1...)
			}
		}
	}

	return userName, people, userImg, arrIndex, arrUp, arrDown, arrDetails
}

//随机不重复
//生成count个(start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int64 {
	//范围检查
	if end < start || (end-start+1) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int64, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start + 1
		//查重
		exist := false
		for _, v := range nums {
			if v == int64(num) {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, int64(num))
		}
	}
	return nums

}

//首页-动态展示
func (b BaseController) addIndex(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
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
	//种类
	var stickArr []StickAll
	for j := 1; j < 8; j++ {

		StickData, err := config.ShowStick(client, big.NewInt(int64(j)))
		if err != nil {
			respError(c, err)
			return
		}
		stickArr1 := []StickAll{StickAll{Sticks: StickData.Name, SticksImg: StickData.Img}}
		stickArr = append(stickArr, stickArr1...)
	}
	fmt.Println(stickArr)
	//定义一个结构体数组

	var goodArr []GoodsPort
	var goodArr1 []GoodsPort
	nums := generateRandomNumber(0, len(id), len(id))
	//轮播
	for j := 0; j < 3; j++ {

		goodData, goodData1, err2 := config.HaveIndex(client, id[j])
		//print("图片路径",goodImg)
		if err2 != nil {
			//respError(c, err2)
			fmt.Println(err2)
			return
		}
		if goodData.Name != "" {
			arr0 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
			goodArr1 = append(goodArr1, arr0...)
			fmt.Println("goodArr", goodArr1)
		}
	}
	for j := 0; j < 3; j++ {
		x := int64(rand.Intn(len(id)))
		goodData, goodData1, err2 := config.HaveIndex(client, big.NewInt(x))
		//print("图片路径",goodImg)
		if err2 != nil {
			fmt.Println("err2", err2)

			return
		}
		if goodData.Name != "" && goodData.Available == true && goodData.IsBorrow == false {
			arr2 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
			goodArr = append(goodArr, arr2...)
		}
	}
	userName, people, userImg, arr1, _, _, _ := b.dynamicHtmlFun(c, len(id), nums, "", "index")

	if err != nil {
		fmt.Println(err)
		return
	}
	c.HTML(http.StatusOK, "Static/index.html", gin.H{
		"goodArr":  goodArr,
		"goodArr1": goodArr1,
		"goodsPor": arr1,
		"userName": userName,
		"userImg":  userImg,
		"address":  people,
		"stickArr": stickArr,
	})
}

//}
//}

//商品详情页展示
func (b BaseController) shopPorduct(c *gin.Context) {
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
	type OrderDet struct {
	}
	goodData, goodData1, err := config.HaveIndex(client, id)
	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	userImg, err := contract.GetUserImg(nil, LoginUser)
	ownerName, _, _, _, _, err := config.GetUserMethod(contract, goodData.Owner)
	ownerImg, err := contract.GetUserImg(nil, goodData1.Owner)

	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/shop-product.html", gin.H{
		"goodsPort": GoodsPort{
			Id:        goodData.Id,
			Names:     goodData.Name,
			Rent:      goodData.Rent,
			EthPledge: goodData.EthPledge,
			GoodImg:   goodData1.GoodImg,
			GoodSign:  goodData1.GoodSign,
			Addr:      goodData.Owner,
			Species:   goodData.Species,
		},
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"ownerName": ownerName,
		"ownerImg":  ownerImg,
	})
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
	for j := 0; j < 9; j++ {
		if j < 8 {
			StickData, err := config.ShowStick(client, big.NewInt(int64(j)))
			if err != nil {
				respError(c, err)
				return
			}
			stickArr1 := []StickAll{StickAll{Sticks: StickData.Name, SticksImg: StickData.Img}}
			stickArr = append(stickArr, stickArr1...)
			for i := 0; i < len(id); i++ {
				goodData, goodData1, _ := config.HaveIndex(client, id[i])
				if goodData.Species == StickData.Name {

					arr1 := []StickAll{StickAll{GoodsPortStick: GoodsPort{Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}}
					stickArr = append(stickArr, arr1...)

				}
			}
		} else {
			//fmt.Println("stickArr===", stickArr)
			//fmt.Println("ashvdlasjvhlja", additionGoods)
			stickArr = append(stickArr, arr...)
			c.HTML(http.StatusOK, "Static/goods-category.html", gin.H{
				"goodsCategory": stickArr,
				"userName":      userName,
				"address":       people,
				"userImg":       userImg,
				//"LimitAdd":      additionGoods,
			})

		}
	}
}

//我的交易-上下架
func (b BaseController) CartGood(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
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
	s := make([]int64, len(id))
	for i := 0; i < len(id); i++ {
		x := id[i].String()
		s[i], err = strconv.ParseInt(x, 10, 64)
	}
	userName, people, userImg, _, arrUp, arrDown, _ := b.dynamicHtmlFun(c, len(id), s, "", "cart")
	c.HTML(http.StatusOK, "Static/cart.html", gin.H{
		"goodsUp":  arrUp,
		"goodDown": arrDown,
		"userName": userName,
		"address":  people,
		"userImg":  userImg,
	})
}

//我借出页
func (b BaseController) MyOrder(c *gin.Context) {
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
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
	s := make([]int64, len(id))
	for i := 0; i < len(id); i++ {
		x := id[i].String()
		s[i], err = strconv.ParseInt(x, 10, 64)
	}
	userName, people, userImg, _, arrUp, _, _ := b.dynamicHtmlFun(c, len(id), s, "", "order")
	//contract, err := config.GetAddress(client)
	//if err != nil {
	//	respError(c, err)
	//	return
	//}
	//userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	//userImg, err := contract.GetUserImg(nil, LoginUser)
	//fmt.Println("res", userName)
	//if err != nil {
	//	respError(c, err)
	//	return
	//}

	//orderId := config.HaveOrderId(client)
	//定义一个结构体数组
	//var arrUp []GoodsPort
	//var arrDown []GoodsPort

	//for i := 0; i < len(orderId)+1; i++ {
	//	fmt.Println("len(orderid)", len(orderId))
	//	if i < len(orderId) {
	//		OrderDet, _ := config.HaveOrderDit(client, orderId[i])
	//		goodData, goodData1, err := config.HaveIndex(client, OrderDet.Id)
	//		isBack, endTime, err := contract.GetBackRec(nil, OrderDet.Id, OrderDet.Back)
	//
	//		fmt.Println("goodData", goodData)
	//		if err != nil {
	//			respError(c, err)
	//			return
	//		}
	//		fmt.Println("orderID", orderId, OrderDet.Id, OrderDet.OId, goodData.Owner, OrderDet.OrderOwner)
	//		if people == OrderDet.OrderOwner {
	//			arr1 := []GoodsPort{GoodsPort{IsBack: isBack, EndTime: endTime, OId: OrderDet.OId, Id: goodData.Id, Names: goodData.Name, Species: goodData.Species, Rent: goodData.Rent, EthPledge: goodData.EthPledge, IsBorrow: goodData.IsBorrow, GoodImg: goodData1.GoodImg, GoodSign: goodData1.GoodSign}}
	//			arrUp = append(arrUp, arr1...)
	//			fmt.Println("arrUp", arrUp)
	//
	//		}
	//
	//	} else {

	c.HTML(http.StatusOK, "Static/order.html", gin.H{
		"goodsUp": arrUp,
		//"goodDown": arrDown,
		"userName": userName,
		"address":  people,
		"userImg":  userImg,
	})
}

//}
//}

//分类详情页面
func (b BaseController) CategoryDetails(c *gin.Context) {
	Spe := c.Params
	Speci := Spe.ByName("Species")
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}

	//contract, err := config.GetAddress(client)
	//if err != nil {
	//	respError(c, err)
	//	return
	//}
	//userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	//userImg, err := contract.GetUserImg(nil, LoginUser)
	//fmt.Println("res", userName)
	//if err != nil {
	//	respError(c, err)
	//	return
	//}
	//获取id
	id := config.HaveId(client)
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	s := make([]int64, len(id))
	for i := 0; i < len(id); i++ {
		x := id[i].String()
		s[i], err = strconv.ParseInt(x, 10, 64)
	}
	userName, people, userImg, _, _, _, arrDetails := b.dynamicHtmlFun(c, len(id), s, Speci, "categoryDetails")
	//var arrDetails []GoodsPort
	//for i := 0; i < len(id)+1; i++ {
	//	if i < len(id) {
	//		//var addr common.Address
	//		goodData, goodData1, err := config.HaveIndex(client, id[i])
	//		if err != nil {
	//			respError(c, err)
	//			return
	//		}
	//
	//		if goodData.Species == Speci {
	//			arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, GoodImg: goodData1.GoodImg}}
	//			arrDetails = append(arrDetails, arr1...)
	//		}
	//		print(arrDetails)
	//	} else {

	c.HTML(http.StatusOK, "Static/category-details.html", gin.H{
		"categoryDetail": arrDetails,
		"userName":       userName,
		"address":        people,
		"userImg":        userImg,
	})
}

//}
//}

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
	orderId := config.HaveOrderId(client)
	var myshoparr []GoodsPort
	var hh1Str string
	//var numx int
	//if len(orderId)>=len(id){
	//	numx=len(orderId)
	//}else{
	//	numx=len(id)
	//}
	for i := 0; i < len(orderId)+1; i++ {
		fmt.Println("len(orderid)", len(orderId))
		if i < len(orderId) {
			OrderDtBorr, _ := config.HaveOrderDit(client, orderId[i])
			goodData, goodData1, err := config.HaveIndex(client, OrderDtBorr.Id)
			_, _, _, days, start, err := contract.GetDealRec(nil, OrderDtBorr.Id, OrderDtBorr.Deal)
			isBack, endTime, err := contract.GetBackRec(nil, OrderDtBorr.Id, OrderDtBorr.Back)
			fmt.Println("orderBrrID", orderId, OrderDtBorr.Id, OrderDtBorr.OId, goodData.Owner, OrderDtBorr.OrderOwner)
			dayStr := days.String()
			daysInt, _ := strconv.Atoi(dayStr)
			hour := daysInt * 24
			day := fmt.Sprintf("%dh", hour)
			hh, _ := time.ParseDuration(day)
			startTime, err := time.Parse("2006-01-02 15:04:05", start)
			hh1 := startTime.Add(hh)
			hh1Str = hh1.Format("2006-01-02 15:04:05")
			fmt.Println("hh1", hh1)
			fmt.Println("myahopar", myshoparr)
			if err != nil {
				respError(c, err)
				return
			}

			if OrderDtBorr.OrderBorrower == people {
				//x, _ := strconv.Atoi(days.String())
				y, _ := strconv.Atoi(goodData.Rent.String())
				z, _ := strconv.Atoi(goodData.EthPledge.String())
				allPrice := daysInt*y + z
				arr1 := []GoodsPort{GoodsPort{Days: daysInt, AllPrice: allPrice, IsBack: isBack, EndTime: endTime, BackTime: hh1Str, OId: OrderDtBorr.OId, Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, EthPledge: goodData.EthPledge, GoodImg: goodData1.GoodImg, Species: goodData.Species}}
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
