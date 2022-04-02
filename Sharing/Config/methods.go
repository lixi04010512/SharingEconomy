package config

import (
	"Sharing/Agreement"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"net/smtp"
	"path"
	"strings"
)

const (
	//chainID = 8565 //8888
	Prikey       = "9e64387a398fa1a813e2a8614cd2ebd04751755d1c2046cb0cecf0498a78591f"
	ShareFishAddress = "0x687168Bf096C205323936E2a435F99D38fEb5676"
	gasLimit      = 3000000
)

//获取client
func GetClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	//client, err := ethclient.Dial("https://rinkeby.etherscan.io")

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	return client, nil
}

// 获取合约对象
func GetAddress(client *ethclient.Client) (*Agreement.User, error) {
	contract, err := Agreement.NewUser(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("faucet:", contract)
	return contract, err
}

//获取Opts对象
func Getopts() *bind.TransactOpts {
	privateKey := GetPrivateKey()
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}
//只读实例 获取UserCaller对象
func HaveUserRead(client *ethclient.Client) (*Agreement.UserCaller, error) {
	contract, err :=Agreement.NewUserCaller(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return contract, err
}

//获取私钥
func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(Prikey)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return big.NewInt(0), err
	} else {
		fmt.Println("gasPrice", gasPrice)
		return gasPrice, nil
	}

}

func GetMsgOpts(privateKey *ecdsa.PrivateKey) *bind.TransactOpts  {
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}

//封装注册方法
func RegisterMethod(client *ethclient.Client,contract *Agreement.User ,Address common.Address,name string,email string,password string) (*types.Transaction, error) {
	opts :=Getopts()
	opts.Value=big.NewInt(1000000000000000000)
	res, err := contract.Register(opts,name,Address,email,password)
	fmt.Println("register:",res)
	opts.GasLimit=3000000
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//发送邮箱验证码方法
func EmailSend(userEmail string)  string{
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "1784420499@qq.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{userEmail}

	// 设置主题
	em.Subject = "ShareFish注册账户"

	code:=EncodeToString(6)
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("[ShareFish],您正在验证邮箱，您的验证码是："+code+",请尽快验证。请勿泄露您的验证码。如非本人操作请忽略。")

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1784420499@qq.com", "qydemwctecedfceb", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("发送成功 ")
	return code
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//封装登录方法
func LoginMethod(client *ethclient.Client,contract *Agreement.User ,Address common.Address,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts:=GetMsgOpts(privateKey)
	res, err := contract.Login(opts,Address)

	fmt.Println("login:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装获取用户信息方法
func GetUserMethod(contract *Agreement.User,address common.Address) (string,common.Address,*big.Int,string,string,*big.Int,*big.Int,error)  {
	res,err:=contract.GetUser(nil,address)
	if err != nil {
		log.Fatal(err)
	}
	return res.Name,res.People,res.Integral,res.Email,res.Sign,res.GoodsNum,res.Balance,nil
}

//封装注销方法
func LogoutMethod(client *ethclient.Client,contract *Agreement.User ,Address common.Address,privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts :=GetMsgOpts(privateKey)
	res, err := contract.Logout(opts,Address)
	fmt.Println("login:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装添加公益方法
func AddCommunityMethod(client *ethclient.Client,contract *Agreement.User ,number *big.Int,beneficiaryAddr common.Address, communityName string,communityIntroduce string,communityAmounts *big.Int) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.AddCommunity(opts,number,beneficiaryAddr,communityName,communityIntroduce,communityAmounts)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装下架公益方法
func DelCommunityMethod(client *ethclient.Client,contract *Agreement.User ,number *big.Int) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.DelCommunity(opts,number)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
//封装捐赠积分方法
func DonateMethod(client *ethclient.Client,contract *Agreement.User ,number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	opts :=Getopts()
	opts.Value=big.NewInt(1000000000000000000)
	res, err := contract.Donate(opts,number,donator,amount,donatorName)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
//封装查看进度方法
func CommunityScheduleView(contract *Agreement.User ,number *big.Int) (*big.Int, error) {
	res, err := contract.CommunitySchedule(nil,number)
	fmt.Println("addCommunity:",res)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品上架方法
func AddGoodsMethod(client *ethclient.Client,contract *Agreement.User ,owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int,goodsImgs []string) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.AddGoods(opts,owner,name,species,rent,ethPledge,goodsImgs)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品下架方法
func DelGoodsMethod(client *ethclient.Client,contract *Agreement.User ,id *big.Int) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.DelGoods(opts,id)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品下架方法
func BorrowGoodsMethod(client *ethclient.Client,contract *Agreement.User , id *big.Int,borrowDays *big.Int) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.BorrowGoods(opts,id,borrowDays)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品归还方法
func DoGoodsReturnMethod(client *ethclient.Client,contract *Agreement.User,id *big.Int) (*types.Transaction, error) {
	opts :=Getopts()
	res, err := contract.DoGoodsReturn(opts,id)
	fmt.Println("addCommunity:",res)
	opts.GasLimit=gasLimit
	opts.GasPrice,err=GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//首页-更多好物
func HaveIndex(client *ethclient.Client ,id *big.Int) (  common.Address,string,string, *big.Int,*big.Int, error) {
	//opts := HaveGetOpts()
	ins, err := HaveUserRead(client)
	//var id *big.Int
	//var opts *bind.CallOpts
	goodsD, err := ins.GetGoods(nil,id)
	fmt.Println("goodid",goodsD)
	if err != nil {
		log.Fatal("GETGOOD错误",err)
	}
	return goodsD.Owner,goodsD.Name,goodsD.Species,goodsD.Rent,goodsD.EthPledge, err
}

//获取id
func HaveId(client *ethclient.Client)[] *big.Int{
	ins,_:=HaveUserRead(client)
	id,err:=ins.GetGoodsId(nil)
	if err != nil {
		log.Fatal("GETGOOD错误",err)
	}
	return id
}
//首页-口碑推荐
func HaveIndex1(client *ethclient.Client ,id *big.Int) ( common.Address,string,string, *big.Int,*big.Int, error) {
	//opts := HaveGetOpts()
	ins, err := HaveUserRead(client)
	//var id *big.Int
	//var opts *bind.CallOpts
	goodsD, err := ins.GetGoods(nil,id)
	fmt.Println("goodid",goodsD)
	if err != nil {
		log.Fatal("GETGOOD错误",err)
	}
	return goodsD.Owner,goodsD.Name,goodsD.Species,goodsD.Rent,goodsD.EthPledge, err
}

//图片上传
func UploadUserImg(c *gin.Context){
	f, err := c.FormFile("imgname")
	//file, handler, err := c.FormFile("imgname")
	fmt.Println(err)
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	} else {
		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".jpeg"{
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,jpeg文件",
			})
			return
		}
		fileName:=f.Filename
		fildDir:="./Static/images"

		filepath:=fmt.Sprintf("%s%s",fildDir,fileName)
		c.SaveUploadedFile(f, filepath)
		fmt.Println("filename",fileName)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result":gin.H{
				"path":filepath,
			},
		})
	}
}