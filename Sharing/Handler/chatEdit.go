package handler

import (
	"Sharing/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查询所有消息列表
var chat []db.Chat_list

func select_chat_list(c *gin.Context) {
	todo, err := db.Select_chat_list()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	fmt.Println(todo)
	chat = todo
	fmt.Println(chat)
}

var addr_owner = LoginUser.Hex()

//消息列表页面
func apps_chat(c *gin.Context) {
	c.HTML(http.StatusOK, "Static/apps-chat.html", gin.H{
		"chat_list": chat,
		"owners":    addr_owner,
	})
}

//通过对方地址查找消息详情
var chat1 []db.Chat

func select_chat_content(c *gin.Context) {
	todo, err := db.Select_chat_content()
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	fmt.Println(todo)
	chat1 = todo
	fmt.Println(chat1)
}

//消息列表跳转消息详情
var addr_x string
var name_x string

func list_content(c *gin.Context) {
	addr := c.PostForm("addr")
	name := c.PostForm("name")
	addr = string([]byte(addr)[:42])
	addr_x = addr
	name_x = name
}

//消息详情页面
func wechat(c *gin.Context) {
	c.HTML(http.StatusOK, "Static/wechat.html", gin.H{
		"chat":   chat1,
		"owners": addr_owner,
		"addrs":  addr_x,
		"names":  name_x,
	})
}

//增加消息列表
func insert_chat_list(c *gin.Context) {
	todo := db.Chat_list{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err := todo.Insert_chat_list()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, "ok")
}

//增加消息详情
func insert_chat_content(c *gin.Context) {
	todo := db.Chat{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err := todo.Insert_chat_content()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, "ok")
}

//删除消息列表
func delete_chat_list(c *gin.Context) {
	todo := db.Chat_list{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err := todo.Delete_chat_list()
	if err != nil {
		respError(c, err)
		return
	}
}

//删除消息详情
func delete_chat_content(c *gin.Context) {
	todo := db.Chat{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err := todo.Delete_chat_content()
	if err != nil {
		respError(c, err)
		return
	}
}
