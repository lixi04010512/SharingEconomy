package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	userName, people, integral, email, sign, goodsNum, balance, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	fmt.Println("peo", people, integral)
	if err != nil {
		respError(c, err)
		return
	}
	c.HTML(http.StatusOK, "Static/profile.html", gin.H{"userName": userName, "address": people, "integral": integral, "email": email, "sign": sign, "goodsNum": goodsNum, "balance": balance})
}

//渲染index
//func IndexStatic(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		fmt.Println(err)
//		respError(c,err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	userName,people,_,_,_,_,_,err:=config.GetUserMethod(contract,loginUser)
//	fmt.Println("res",userName)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//	c.HTML(http.StatusOK, "Static/index.html", gin.H{"userName": userName,"address": people})
//}

////渲染cart
//func CartStatic(c *gin.Context) {
//	//初始化client
//	client, err := config.GetClient()
//	if err != nil {
//		fmt.Println(err)
//		respError(c, err)
//		return
//	}
//	//初始化合约地址
//	contract, err := config.GetAddress(client)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//
//	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
//	fmt.Println("res", userName)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//	userImg,err :=contract.GetUserImg(nil,loginUser)
//	c.HTML(http.StatusOK, "Static/cart.html", gin.H{"userName": userName, "address": people,"userImg":userImg})
//}

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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/category-details.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

//渲染chat
func ChatStatic(c *gin.Context) {
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/chat.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/edit-profile.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

////渲染goods-category
//func GoodsStatic(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		fmt.Println(err)
//		respError(c,err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	userName,people,_,_,_,_,_,err:=config.GetUserMethod(contract,loginUser)
//	fmt.Println("res",userName)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//	c.HTML(http.StatusOK, "Static/goods-category.html", gin.H{"userName": userName,"address": people})
//}

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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/lend-items.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/post-need.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/shop.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}

////渲染shop-profile
//func ShopProfilerStatic(c *gin.Context)  {
//	//初始化client
//	client,err := config.GetClient()
//	if err != nil{
//		fmt.Println(err)
//		respError(c,err)
//		return
//	}
//	//初始化合约地址
//	contract ,err:= config.GetAddress(client)
//	if err != nil{
//		respError(c,err)
//		return
//	}
//
//	userName,people,_,_,_,_,_,err:=config.GetUserMethod(contract,loginUser)
//	fmt.Println("res",userName)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//	c.HTML(http.StatusOK, "Static/shop-product.html", gin.H{"userName": userName,"address": people})
//}

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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
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

	userName, people, _, _, _, _, _, err := config.GetUserMethod(contract, loginUser)
	fmt.Println("res", userName)
	if err != nil {
		respError(c, err)
		return
	}
	userImg, err := contract.GetUserImg(nil, loginUser)
	c.HTML(http.StatusOK, "Static/wishlist.html", gin.H{"userName": userName, "address": people, "userImg": userImg})
}
