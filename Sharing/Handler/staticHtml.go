package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

//获取用户信息
func UserProfile(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, email, sign, goodsNum, err := config.GetUserMethod(contract, LoginUser)
	balances, err := contract.GetBalance(nil, LoginUser)

	balanceStr := balances.String()
	balanceF, err := strconv.ParseFloat(balanceStr, 64)
	balance, err := strconv.ParseFloat(fmt.Sprintf("%.2f", balanceF/1000000000000000000), 256)
	fmt.Println("res", userName)
	fmt.Println("peo", people, balances)
	userImg, err := contract.GetUserImg(nil, LoginUser)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/profile.html", gin.H{"userName": userName, "address": people, "email": email, "sign": sign, "goodsNum": goodsNum, "balance": balance, "userImg": userImg})
}

//渲染mod-photo
func modPhotoStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/mod-photo.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染category-details
func CategoryStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/category-details.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染tally-order
func TallyStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	idInt64 := int64(idInt)
	goodsData, goodsData1, err := config.HaveIndex(client, big.NewInt(idInt64))
	c.HTML(http.StatusOK, "Static/tally-order.html", gin.H{
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"goodsImg":  goodsData1.GoodImg,
		"goodsName": goodsData.Name,
		"rent":      goodsData.Rent,
		"pledge":    goodsData.EthPledge,
		"id":        id,
	})
}

//渲染edit-need
func EditNeedStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/edit-need.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染edit-profile
func EditProfileStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, email, sign, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/edit-profile.html", gin.H{"userName": userName, "address": people, "userImg": userImg, "userEmail": email, "userSign": sign})
}

//渲染lend-items
func LendStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	var SpeciesArr []string
	for i := 0; i < 30; i++ {
		Int64 := int64(i)
		name, err := config.ShowSpecies(contract, big.NewInt(Int64))
		if err != nil {
			respError(c, err)
			return
		}
		if name != "" {
			SpeciesArr = append(SpeciesArr, name)
		}

	}
	c.HTML(http.StatusOK, "Static/lend-items.html", gin.H{
		"userName": userName,
		"address":  people,
		"userImg":  userImg,
		"spArr":    SpeciesArr,
	})
}

//渲染login
func LoginStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/login.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染myneed
func MyNeedStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/myneed.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染order
func OrderStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/order.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染post-need
func PostNeedStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	var stickArr []StickAll
	for j := 1; j < 9; j++ {
		//if j < 8 {
		StickData, err := config.ShowStick(client, big.NewInt(int64(j)))
		if err != nil {
			respError(c, err)
			return
		}
		stickArr1 := []StickAll{StickAll{Sticks: StickData.Name, SticksImg: StickData.Img}}
		stickArr = append(stickArr, stickArr1...)
		//}
	}
	userImg, _ := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/post-need.html", gin.H{
		"userName":  userName,
		"address":   people,
		"userImg":   userImg,
		"StickList": stickArr,
	})
}

//渲染public-benefit
func PubBenefitStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/public-benefit.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染public-details
func PubDetailsStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/public-details.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染register
func RegisterStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/register.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染search-page
func SearchPageStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/search-result.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染shop
func ShopStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/shop.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染ui-me
func UiMeStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)

	c.HTML(http.StatusOK, "Static/ui-me.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染wishlist
func WishlistStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/wishlist.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染pay
func PayStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/pay.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染con
func ConfirmStatic(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}

	userName, people, _, _, _, err := config.GetUserMethod(contract, LoginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, LoginUser)
	c.HTML(http.StatusOK, "Static/confirm-transaction.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}
