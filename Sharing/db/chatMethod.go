package db

import (
	"fmt"
	"time"
)

//查询消息列表
func Select_chat_list() (todos []Chat_list, err error) {
	var sqlStr = `select * from chat_list`

	r, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("35:", err)
		return
	}
	for r.Next() {
		todo := Chat_list{}
		err = r.Scan(&todo.Id, &todo.Owner, &todo.Addr, &todo.Name, &todo.Img, &todo.Time)
		if err != nil {
			fmt.Println("42:", err)
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select chat_list succ:", todos)
	return todos, err
}

//查询消息详情
func Select_chat_content() (todos []Chat, err error) {
	var sqlStr = `select * from chat `

	r, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("65:", err)
		return
	}
	for r.Next() {
		todo := Chat{}
		err = r.Scan(&todo.Id, &todo.Addr_from, &todo.Addr_to, &todo.Content, &todo.Img_from, &todo.Img_to, &todo.Time)
		if err != nil {
			fmt.Println("64:", err)
			return
		}
		todos = append(todos, todo)
	}
	fmt.Println("select chat succ:", todos)
	return todos, err
}

//增加消息列表
func (chat *Chat_list) Insert_chat_list() (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	fmt.Println("all:",chat.Owner, chat.Addr, chat.Name, chat.Img)
	r, err := db.Exec("insert into chat_list(owner,addr,name,img,time)values(?, ?,?,?,?)", chat.Owner, chat.Addr, chat.Name, chat.Img, strTime)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert chat_list succ:", id)
	return
}

//增加消息详情
func (chat *Chat) Insert_chat_content() (err error) {
	times := time.Now().Local()
	strTime := times.Format("15:05:05")
	r, err := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", chat.Addr_from, chat.Addr_to, chat.Content, chat.Img_from, chat.Img_to, strTime)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert chat succ:", id)
	return
}

//删除消息列表
func (chat *Chat_list) Delete_chat_list() (err error) {
	fmt.Println("chatlist:", chat.Id)
	r, err := db.Exec("delete from chat_list where id=?", chat.Id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("delete chat_list succ:", r)
	return
}

//删除消息详情
func (chat *Chat) Delete_chat_content() (err error) {
	fmt.Println("chat:", chat.Id)
	r, err := db.Exec("delete from chat where id=? ", chat.Id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete chat succ:", r)
	return
}
