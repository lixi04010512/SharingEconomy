package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
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
	r.POST("/addCommunities", getting, addCommunities)
	r.GET("/species", getting, showSpecies)
	r.GET("/showCommunities", getting, showCommunities)
	r.GET("/user", getting, user)
	r.GET("/addSpecies", getting, addSpecies)
	r.GET("/addCommunity", getting, addCommunity)
	r.POST("/editSpecies", getting, editSpecies)
	r.GET("/editCommunity", getting, editCommunity)
	r.POST("/editCommunities", getting, editCommunities)
	r.POST("/updateCommunity", getting, updateCommunity)
	r.POST("/delSpecie", getting, delSpecie)
	r.POST("/delCommunity", getting, delCommunity)
	r.GET("/editSpeciesPage", getting, editSpeciesPage)
	r.POST("/updateStick", getting, updateStick)

	// api接口服务，定义了路由组 /Sharing
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
	}
	share := r.Group("")
	{
		share.GET("/profile", UserProfile)
		share.GET("/index", addIndex)
		share.GET("/cart", CartGood)
		share.GET("/category-details", CategoryDetails)
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
		share.GET("/shop-product", shopPorduct)
		share.GET("/ui-me", UiMeStatic)
		share.GET("/wishlist", WishlistStatic)
		share.GET("/carts", CartGood)
		share.GET("/mod-photo", modPhotoStatic)

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
