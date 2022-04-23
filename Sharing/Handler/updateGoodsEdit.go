package handler

//func updateGoods(c *gin.Context) {
//	//初始化client
//	client, err := config.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	//初始化合约地址
//	contract, err := config.GetAddress(client)
//	if err != nil {
//		respError(c, err)
//		return
//	}
//
//	id := c.PostForm("delId")
//	idInt, err := strconv.Atoi(id)
//	idInt64 := int64(idInt)
//	_, goodData1, err := config.HaveIndex(client, big.NewInt(idInt64))
//	data,err := config.UpdateGoodsMethod(client,contract,big.NewInt(idInt64),goodData1.Name,goodData1.Species,goodData1.Rent,goodData1.EthPledge,privKey)
//	fmt.Println("data",data)
//}