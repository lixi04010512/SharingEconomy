package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		//"addGoodsCategory": addGoodsCategory,
	})
	r.LoadHTMLGlob("Static/*.html")
	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器
		r.Static("/share", webDir)
	}

	// 基于cookie创建session的存储引擎，传递一个参数，用来做加密时的密钥
	store := cookie.NewStore([]byte("secret1234"))
	fmt.Println(store)
	//session中间件生效，参数mysession，是浏览器端cookie的名字
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/manager-login", login1)
	r.POST("/set", setting, loginManager)
	r.GET("/index1", getting, index1)
	r.POST("/addSticks", getting, addSticks)
	r.GET("/species", getting, showSpecies)
	r.GET("/user", getting, user)
	r.GET("/addSpecies", getting, addSpecies)
	r.POST("/editSpecies", getting, editSpecies)
	r.POST("/delSpecie", getting, delSpecie)
	r.GET("/editSpeciesPage", getting, editSpeciesPage)
	r.POST("/updateStick", getting, updateStick)

	r.GET("/select_chat_list", select_chat_list, apps_chat)
	r.POST("/list_content", list_content)
	r.GET("/select_chat_content", select_chat_content, wechat)
	r.POST("/insert_chat_list", insert_chat_list)
	r.POST("/insert_chat_content", insert_chat_content)
	r.POST("/delete_chat_list", delete_chat_list)
	r.POST("/delete_chat_content", delete_chat_content)

	// api接口服务，定义了路由组
	todo := r.Group("")
	{
		// 定义增改查的接口，并注册到web服务器

		//todo.POST("/addStickerPost", addSticker)
		//todo.POST("/isStickExistPost", isStickExist)
		todo.POST("/registerPost", register)
		todo.POST("/loginPost", login)
		todo.POST("/logoutPost", logout)
		todo.POST("/sendEmailPost", sendEmail)
		todo.POST("/addUserImg", AddUserImg)
		todo.POST("/addGoodsPost", addGoods)
		todo.POST("/priLoginPost", privateLogin)
		todo.POST("/updateUserPost", updateUser)
		todo.POST("/addDemandPost", demandAdd)
		todo.POST("/MyListNeeds", MYList)
		todo.POST("/ListNeeds", ListNeedAll)
		todo.GET("/delNeeds/:id", delNeeds)
		todo.POST("/delGoodsPost", delGoods)
		todo.POST("/topUpPost", topUp)
		todo.POST("/sendBorrowMsg", BorrowGoods)
		todo.POST("/searchGoods", searchGoods)
	}
	share := r.Group("")
	{
		share.GET("/profile", UserProfile)
		share.GET("/index", addIndex)
		share.GET("/cart", CartGood)
		share.GET("/category-details/:Species", CategoryDetails)
		share.GET("/chat", ChatStatic)
		share.GET("/edit-need", EditNeedStatic)
		share.GET("/edit-profile", EditProfileStatic)
		share.GET("/goods-category", goodsCategory)
		share.GET("/lend-items", LendStatic)
		share.GET("/login", LoginStatic)
		share.GET("/myneed", MyNeedStatic)
		share.GET("/order", MyOrder)
		share.GET("/post-need", PostNeedStatic)
		share.GET("/public-benefit", PubBenefitStatic)
		share.GET("/public-details", PubDetailsStatic)
		share.GET("/register", RegisterStatic)
		share.GET("/search-page", SearchPageStatic)
		share.GET("/shop", Myshop)
		share.GET("/shop-product/:id", shopPorduct)
		share.GET("/ui-me", UiMeStatic)
		share.GET("/wishlist", WishlistStatic)
		share.GET("/carts", CartGood)
		share.GET("/mod-photo", modPhotoStatic)
		share.GET("/order-receiving/:id", orderReceiving)
		share.GET("/pay", PayStatic)
		share.GET("/tally-order/:id", TallyStatic)
		share.GET("/confirm-transaction", ConfirmStatic)

	}
	// 启动web服务
	err = r.Run(addr)
	return err
}

// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
