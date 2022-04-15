package config

import (
	"Sharing/Agreement"
	"context"
	"crypto/ecdsa"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	//chainID = 8565 //8888
	Prikey       = "5fd7eb82fbd4c5cc87fbeb867ff4804537723d80a5289606c28ca5c8eed4aa77"
	ShareFishAddress = "0xC6b698151b4B257dFF4883A8E43e59B8425e6d08"
	//Prikey       = "9e64387a398fa1a813e2a8614cd2ebd04751755d1c2046cb0cecf0498a78591f"
	//ShareFishAddress = "0x2c49b8475Af3CBC2b64C491AF5a0C761700aD1E7"
	gasLimit = 3000000
	//fileKeystore     = "UTC--2022-03-17T08-08-40.600466800Z--59b0f8a34d8f0dd0e0eef44d02cef0c12fffb9de"
	//Prikey = "c5b9c7fd467335bd829b3b2a3098a72ac39b7f5efa162220b7907cfc684df9a3"
	//privateKey       = "111"
	//ShareFishAddress = "0x1415f8284C54Fbbf5b5300B6177a24A491b1bd07"
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
func HaveClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chainID)
	return client, err
}

//只写实例 获取UserCaller对象
func HaveUserWrite(client *ethclient.Client) (*Agreement.UserTransactor, error) {
	ins, err := Agreement.NewUserTransactor(common.HexToAddress(ShareFishAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return ins, err
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
	contract, err := Agreement.NewUserCaller(common.HexToAddress(ShareFishAddress), client)
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

//获取账户文件和密码 opts
//func HaveGetOpts() *bind.TransactOpts {
//	//账户文件
//	b, err := ioutil.ReadFile(fileKeystore)
//	if err != nil {
//		log.Fatal("错误", err)
//	}
//	//账户密码
//	key, err := keystore.DecryptKey(b, privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("key:", key)
//	reader, _ := os.Open(fileKeystore)
//	//opts, err := bind.NewTransactor(reader, "111")
//	opts, err := bind.NewTransactorWithChainID(reader, "111", big.NewInt(9696))
//	if err != nil {
//		log.Fatal("NewTransactor ", err)
//	}
//	return opts
//}

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

func GetMsgOpts(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(privateKey)
	fmt.Println("opts:", opts)
	return opts
}

<<<<<<< HEAD
=======
//封装注册方法
func RegisterMethod(client *ethclient.Client, contract *Agreement.User, Address common.Address, name string, email string, password string) (*types.Transaction, error) {
	opts := Getopts()
	opts.Value = big.NewInt(1000000000000000000)
	res, err := contract.Register(opts, name, Address, email, password)
	fmt.Println("register:", res)
	opts.GasLimit = 3000000
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//发送邮箱验证码方法
func EmailSend(userEmail string) string {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "1784420499@qq.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{userEmail}

	// 设置主题
	em.Subject = "ShareFish注册账户"

	code := EncodeToString(6)
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("[ShareFish],您正在验证邮箱，您的验证码是：" + code + ",请尽快验证。请勿泄露您的验证码。如非本人操作请忽略。")

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
func LoginMethod(client *ethclient.Client, contract *Agreement.User, Address common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.Login(opts, Address)

	fmt.Println("login:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装获取用户信息方法
func GetUserMethod(contract *Agreement.User, address common.Address) (string, common.Address, *big.Int, string, string, *big.Int, *big.Int, error) {
	res, err := contract.GetUser(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return res.Name, res.People, res.Integral, res.Email, res.Sign, res.GoodsNum, res.Balance, nil
}

//封装注销方法
func LogoutMethod(client *ethclient.Client, contract *Agreement.User, Address common.Address, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	opts := GetMsgOpts(privateKey)
	res, err := contract.Logout(opts, Address)
	fmt.Println("login:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装捐赠积分方法
func DonateMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	opts := Getopts()
	opts.Value = big.NewInt(1000000000000000000)
	res, err := contract.Donate(opts, number, donator, amount, donatorName)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品上架方法
func AddGoodsMethod(client *ethclient.Client, contract *Agreement.User, owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string, goodsSign string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddGoods(opts, owner, name, species, rent, ethPledge, goodsImgs, goodsSign)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品下架方法
func DelGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DelGoods(opts, id)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品下架方法
func BorrowGoodsMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.BorrowGoods(opts, id, borrowDays)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装物品归还方法
func DoGoodsReturnMethod(client *ethclient.Client, contract *Agreement.User, id *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DoGoodsReturn(opts, id)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//首页-更多好物
func HaveIndex(client *ethclient.Client, id *big.Int) (
	struct {
		Owner     common.Address
		Borrowers Agreement.UserBorrower
		Id        *big.Int
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
		Count     *big.Int
		Deal      *big.Int
		Backs     *big.Int
		Available bool
		IsBorrow  bool
	}, struct {
		Owner     common.Address
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
		GoodImg   []string
	}, error) {
	ins, err := HaveUserRead(client)
	goodsD, err := ins.GoodsData(nil, id)
	if err != nil {
		log.Fatal(err)
	}
	goodsD1, err := ins.GetGoods(nil, id)
	return goodsD, goodsD1, err
>>>>>>> e1bad8d2d96ac8b69a502927e4206259c58cab65



<<<<<<< HEAD

=======
//图片上传
func UploadUserImg(c *gin.Context) (string, error) {
	f, err := c.FormFile("imgname")
	fileName := f.Filename
	fildDir := "./Static/images"

	filePath := fmt.Sprintf("%s%s", fildDir, fileName)
	//file, handler, err := c.FormFile("imgname")
	fmt.Println(err)
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,jpeg文件",
			})
			return "", err
		}

		c.SaveUploadedFile(f, filePath)
		fmt.Println("filename", fileName)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result": gin.H{
				"path": filePath,
			},
		})
	}
	return filePath, err
}

//封装头像上传方法
func AddUserImgMethod(client *ethclient.Client, contract *Agreement.User, people common.Address, headImg string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddUserImg(opts, people, headImg)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装管理员登录方法
func LoginManager(contract *Agreement.User, Address common.Address) (string, error) {
	res, err := contract.SignIn(nil, Address)

	fmt.Println("login:", res)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装添加分类方法
func AddStick(client *ethclient.Client, contract *Agreement.User, stick string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddSticker(opts, stick)
	fmt.Println("addstick:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装分类展示方法
func ShowSpecies(contract *Agreement.User, id *big.Int) (string, error) {
	res, err := contract.GetStick(nil, id)

	fmt.Println("showSpecies:", res)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装删除分类方法
func DelSpecieMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DelStick(opts, number)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	fmt.Println("Delstick:", err)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装修改分类方法
func UpdateStickMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int, name string) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.UpdateStick(opts, number, name)
	fmt.Println("updateStick:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装添加公益方法
func AddCommunityMethod(client *ethclient.Client, contract *Agreement.User, beneficiaryAddr common.Address, communityName string, communityIntroduce string, communityAmounts *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.AddCommunity(opts, beneficiaryAddr, communityName, communityIntroduce, communityAmounts)
	fmt.Println("addCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装下架公益方法
func DelCommunityMethod(client *ethclient.Client, contract *Agreement.User, number *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.DelCommunity(opts, number)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	fmt.Println("DelCommunity:", err)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

//封装公益展示方法
func ShowCommunities(contract *Agreement.User, id *big.Int) (string, common.Address, string, *big.Int, error) {
	community, err := contract.GetCommunity(nil, id)
	fmt.Println("showCommunities:", community)
	if err != nil {
		log.Fatal(err)
	}
	return community.Name, community.People, community.Introduce, community.Amount, nil
}

//封装修改公益方法
func UpdateCommunityMethod(client *ethclient.Client, contract *Agreement.User, name string, number *big.Int, addr common.Address, introduce string, amount *big.Int) (*types.Transaction, error) {
	opts := Getopts()
	res, err := contract.UpdateCommunity(opts, number, addr, name, introduce, amount)
	fmt.Println("updateCommunity:", res)
	opts.GasLimit = gasLimit
	opts.GasPrice, err = GetgasPrice(client)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
>>>>>>> e1bad8d2d96ac8b69a502927e4206259c58cab65
