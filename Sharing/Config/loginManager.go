package config

import (
	"Sharing/Agreement"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

//封装管理员登录方法
func LoginManager(contract *Agreement.User, Address common.Address) (string, error) {
	res, err := contract.SignIn(nil, Address)

	fmt.Println("login:", res)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
