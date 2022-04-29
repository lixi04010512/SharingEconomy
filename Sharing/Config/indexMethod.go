package config

import (
	"Sharing/Agreement"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

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
		IsAgree   bool
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

}

//获取id
func HaveId(client *ethclient.Client) []*big.Int {
	ins, _ := HaveUserRead(client)
	id, err := ins.GetGoodsId(nil)
	if err != nil {
		log.Fatal("GETGOOD错误", err)
	}
	return id
}
