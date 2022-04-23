package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
)


var jwtkey = []byte("www.topgoer.com")
var str string

type Claims struct {
	UserPassword string `form:"password"`
	jwt.StandardClaims
}
type Stick struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Community struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	People    common.Address `json:"people"`
	Introduce string         `json:"introduce"`
	Amount    *big.Int       `json:"amount"`
}

//登录页面
func login1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "manager-login.html", gin.H{})
}

//首页
func index1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index1.html", gin.H{})

}

//用户管理页
func user(c *gin.Context) {
	c.HTML(http.StatusOK, "Static/user.html", gin.H{})
}

//添加分类页面
func addSpecies(c *gin.Context) {
	c.HTML(http.StatusOK, "addSpecies.html", gin.H{})
}

//添加公益页面
func addCommunity(c *gin.Context) {
	c.HTML(http.StatusOK, "addCommunity.html", gin.H{})
}

//定义一个结构体数组
var arr []Stick

//修改分类页面
func editSpecies(c *gin.Context) {
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
		fmt.Println("260:", err)
		respError(c, err)
		return
	}

	id := c.PostForm("id")
	fmt.Println(id)
	id_specie, err := strconv.Atoi(id)
	i := int64(id_specie)
	name, err := config.ShowSpecies(contract, big.NewInt(i))
	fmt.Println("270:", err)
	arr1 := []Stick{{Id: id_specie, Name: name}}
	fmt.Println("arr1", arr1)
	var arr2 []Stick
	arr2 = append(arr2, arr1...)
	arr = arr2
	fmt.Println("specie:", name)
	fmt.Println("arr:", arr)

}

//修改分类页面
func editSpeciesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Static/editSpecies.html", gin.H{
		"species_name": arr,
	})

}

////定义一个结构体数组
//var arr1 []Community
//
////修改公益页面
//func editCommunities(c *gin.Context) {
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
//		fmt.Println("304:", err)
//		respError(c, err)
//		return
//	}
//
//	id := c.PostForm("id")
//	fmt.Println(id)
//	id_community, err := strconv.Atoi(id)
//	i := int64(id_community)
//	name, addr, introduce, amount, err := config.ShowCommunities(contract, big.NewInt(i))
//	fmt.Println("137:", err)
//	arr2 := []Community{{Id: id_community, Name: name, People: addr, Introduce: introduce, Amount: amount}}
//	fmt.Println("arr2", arr2)
//	var arr3 []Community
//	arr3 = append(arr3, arr2...)
//	arr1 = arr3
//	fmt.Println("community:", name)
//	fmt.Println("arr:", arr)
//
//}

////修改公益页面
//func editCommunity(c *gin.Context) {
//	c.HTML(http.StatusOK, "Static/editCommunity.html", gin.H{
//		"communities_name": arr1,
//	})
//}

//颁发token
func setting(ctx *gin.Context) {
	password := ctx.PostForm("login-password")
	expireTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		UserPassword: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println("351:", err)
	}
	str = tokenString
	session := sessions.Default(ctx)
	session.Set("token", tokenString)
	session.Save()
}

//解析token
func getting(ctx *gin.Context) {
	session := sessions.Default(ctx)
	tokenString := session.Get("token")
	fmt.Println("tokenString:", tokenString)
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "为空"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString.(string))
	if err != nil || !token.Valid {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "出错"})
		ctx.Abort()
		return
	}

	fmt.Println(claims)
	fmt.Println(claims.UserPassword)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		fmt.Println(err)
		return jwtkey, nil
	})
	return token, Claims, err
}

//管理员登录
func loginManager(c *gin.Context) {
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
		fmt.Println("101:", err)
		respError(c, err)
		return
	}

	_, f, err := c.Request.FormFile("login-username")
	if err != nil {
		fmt.Println("410:", err)
		respError(c, err)
		return
	}
	password := c.PostForm("login-password")
	fmt.Println(password)

	keyjson, errs := ioutil.ReadFile("./keystore/" + f.Filename)
	fmt.Println("keyjson:", string(keyjson))
	if errs != nil {
		respError(c, err)
		return
	}
	unlockedKey, errors := keystore.DecryptKey(keyjson, password)
	fmt.Println("unlock:", unlockedKey)
	if errors != nil {
		respError(c, err)
		fmt.Println("480:", errors)
		return
	}
	comAddr := unlockedKey.Address
	data, err := config.LoginManager(contract, comAddr)
	fmt.Println("485err:", err)
	addr := comAddr.Hex()

	c.Redirect(http.StatusFound, "/index1")
	//respOK(c,data)
	fmt.Println("addr", addr)
	fmt.Println("data", data)
}

//添加分类
func addSticks(c *gin.Context) {
	//初始化client
	client, err := config.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	//初始化合约地址
	contract, err := config.GetAddress(client)
	if err != nil {
		respError(c, err)
		return
	}
	species := c.PostForm("species_name")
	data, err := config.AddStick(client, contract, species)

	fmt.Println("sticks", data)
}

//分类动态展示
func showSpecies(c *gin.Context) {
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
		fmt.Println("472:", err)
		respError(c, err)
		return
	}

	//定义一个结构体数组
	var arr []Stick
	for i := 1; i < 25; i++ {
		//转*big.int
		Int64 := int64(i)
		name, err := config.ShowSpecies(contract, big.NewInt(Int64))
		fmt.Println("483:", err)
		if name != "" {
			arr1 := []Stick{{Id: i, Name: name}}
			fmt.Println("arr1", arr1)
			arr = append(arr, arr1...)
			fmt.Println("species:", name)
			fmt.Println("arr", arr)
		}

	}

	c.HTML(http.StatusOK, "Static/species.html", gin.H{
		"species_name": arr,
	})

}

//删除分类
func delSpecie(c *gin.Context) {
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
		fmt.Println("512:", err)
		respError(c, err)
		return
	}
	id := c.PostForm("id")
	fmt.Println(id)
	id_specie, err := strconv.Atoi(id)
	i := int64(id_specie)
	data, err := config.DelSpecieMethod(client, contract, big.NewInt(i))
	fmt.Println("521data:", data)
}

//修改分类
func updateStick(c *gin.Context) {
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
		fmt.Println("536:", err)
		respError(c, err)
		return
	}
	id := c.PostForm("id")
	id1, err := strconv.Atoi(id)
	id2 := int64(id1)
	name := c.PostForm("name")
	data, err := config.UpdateStickMethod(client, contract, big.NewInt(id2), name)
	fmt.Println(data)
}

////修改公益
//func updateCommunity(c *gin.Context) {
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
//		fmt.Println("560:", err)
//		respError(c, err)
//		return
//	}
//	id := c.PostForm("id")
//	id1, err := strconv.Atoi(id)
//	id2 := int64(id1)
//	name := c.PostForm("name")
//	addr := c.PostForm("addr")
//	var data1 = []byte(addr)
//	address := common.BytesToAddress(data1)
//	introduce := c.PostForm("introduce")
//	amount := c.PostForm("amount")
//	money, err := strconv.Atoi(amount)
//	Money := int64(money)
//	data, err := config.UpdateCommunityMethod(client, contract, name, big.NewInt(id2), address, introduce, big.NewInt(Money))
//	fmt.Println(data)
//}
//
////添加公益
//func addCommunities(c *gin.Context) {
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
//		fmt.Println("591:", err)
//		respError(c, err)
//		return
//	}
//	addr := c.PostForm("addr")
//	var data1 = []byte(addr)
//	address := common.BytesToAddress(data1)
//	name := c.PostForm("name")
//	introduce := c.PostForm("introduce")
//	money := c.PostForm("amount")
//	money1, err := strconv.Atoi(money)
//	Money := int64(money1)
//	fmt.Println(address, name, introduce, Money)
//	data, err := config.AddCommunityMethod(client, contract, address, name, introduce, big.NewInt(Money))
//	fmt.Println("community,", data)
//}
//
////删除公益
//func delCommunity(c *gin.Context) {
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
//		fmt.Println("620:", err)
//		respError(c, err)
//		return
//	}
//	id := c.PostForm("id")
//	fmt.Println(id)
//	id_community, err := strconv.Atoi(id)
//	i := int64(id_community)
//	data, err := config.DelCommunityMethod(client, contract, big.NewInt(i))
//	fmt.Println("629data:", data)
//}
//
////公益动态展示
//func showCommunities(c *gin.Context) {
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
//		fmt.Println("644:", err)
//		respError(c, err)
//		return
//	}
//
//	//定义一个结构体数组
//	var arr []Community
//	for i := 1; i < 10; i++ {
//		//转*big.int
//		Int64 := int64(i)
//		name, people, introduce, amount, err := config.ShowCommunities(contract, big.NewInt(Int64))
//		fmt.Println("655:", err)
//		if name != "" {
//			arr1 := []Community{{Id: i, Name: name, People: people, Introduce: introduce, Amount: amount}}
//			fmt.Println("arr1", arr1)
//			arr = append(arr, arr1...)
//			fmt.Println("community:", name)
//			fmt.Println("arr:", arr)
//		}
//
//	}
//
//	c.HTML(http.StatusOK, "Static/community.html", gin.H{
//		"communities_name": arr,
//	})
//
//}

