package handler

import (
	config "Sharing/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

//模糊匹配 b中找a

func StrVagueQuery(a string, b string) bool {
	return strings.Contains(b, a)

}

func searchGoods(c *gin.Context) {
	keysGoods := c.PostForm("keysGoods")
	client, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	//获取id
	id := config.HaveId(client)
	if err != nil {
		respError(c, err)
		return
	}
	var searchGoods []GoodsPort
	for i := 0; i < len(id); i++ {
		goodData, goodData1, err := config.HaveIndex(client, id[i])
		if err != nil {
			respError(c, err)
			return
		}
		if (goodData.Available == true && goodData.IsBorrow == false) && (StrVagueQuery(keysGoods, goodData.Name) || StrVagueQuery(keysGoods, goodData.Species)) {
			arr1 := []GoodsPort{GoodsPort{Id: goodData.Id, Names: goodData.Name, Rent: goodData.Rent, GoodImg: goodData1.GoodImg, Species: goodData.Species}}
			searchGoods = append(searchGoods, arr1...)
		}
	}
	fmt.Println("searchGoods", searchGoods)
	respOK(c, searchGoods)
}
