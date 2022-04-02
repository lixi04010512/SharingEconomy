package handler

import (
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

	// api接口服务，定义了路由组 /Sharing
	todo := r.Group("")
	{
		// 定义增改查的接口，并注册到web服务器

		todo.POST("/addStickerPost",addSticker)
		todo.POST("/isStickExistPost",isStickExist)
		todo.POST("/registerPost",register)
		todo.POST("/loginPost",login)
		todo.POST("/logoutPost",logout)
		todo.POST("/sendEmailPost",sendEmail)

	}
	share := r.Group("")
	{
		share.GET("/profile",UserProfile)
		share.GET("/index",IndexStatic)
		share.GET("/cart",CartStatic)
		share.GET("/category-details",CategoryStatic)
		share.GET("/chat",ChatStatic)
		share.GET("/edit-need",EditNeedStatic)
		share.GET("/edit-profile",EditProfileStatic)
		share.GET("/goods-category",GoodsStatic)
		share.GET("/lend-items",LendStatic)
		share.GET("/login",LoginStatic)
		share.GET("/myneed",MyNeedStatic)
		share.GET("/order",OrderStatic)
		share.GET("/post-need",PostNeedStatic)
		share.GET("/public-benefit",PubBenefitStatic)
		share.GET("/public-details",PubDetailsStatic)
		share.GET("/register",RegisterStatic)
		share.GET("/search-page",SearchPageStatic)
		share.GET("/shop",ShopStatic)
		share.GET("/shop-product",ShopProfilerStatic)
		share.GET("/ui-me",UiMeStatic)
		share.GET("/wishlist",WishlistStatic)

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

