package config

import (
	"Sharing/Agreement"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jordan-wright/email"
	"io"
	"log"
	"math/big"
	"net/smtp"
)


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

//发送邮箱附件方法
func EmailSend1(userEmail string,name string,key string)  error {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "1784420499@qq.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{userEmail}

	// 设置主题
	em.Subject = "ShareFish注册账户"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("[ShareFish],您已成功注册，您的私钥是："+key+",您的账户文件如下,请不要泄露个人信息。")
	em.AttachFile("./keystore/"+name)

	//设置服务器相关的配置
	err1 := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1784420499@qq.com", "qydemwctecedfceb", "smtp.qq.com"))
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println("发送成功 ")
	return err1
}