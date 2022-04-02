package main

import (
	"Sharing/Handler"
	"fmt"
)

func main() {
	// 初始化web服务器
	err := handler.Start(fmt.Sprintf("%s:%s", "127.0.0.1", "8888"),"Static")
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}

}
