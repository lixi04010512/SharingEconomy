package db

import (
	"fmt"
	"strconv"
	"time"
)

//查询消息列表
func Select_chat_list() (todos []Chat_list, err error) {
	var sqlStr = `select * from chat_list`
	r, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("14:", err)
		return
	}
	for r.Next() {
		todo := Chat_list{}
		err = r.Scan(&todo.Id, &todo.Owner, &todo.Addr, &todo.Name, &todo.Img, &todo.Time, &todo.No_read)
		if err != nil {
			fmt.Println("21:", err)
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

//点开某人的聊天详情，未读归0
func Update_chat_list(addr string, owner string) (err error) {
	r, err := db.Exec("update chat_list set no_read=0 where owner=? and addr=?", owner, addr)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("update chat_list succ:", r)
	return
}

//增加消息列表
var id int

func (chat *Chat_list) Insert_chat_list(username string, userimg string) (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	fmt.Println("all:", chat.Owner, chat.Addr, chat.Name, chat.Img)
	img := "share/" + chat.Img
	userimg = "share/" + userimg
	r1, err := db.Query("select * from chat_list where addr=? and owner=?", chat.Addr, chat.Owner)
	for r1.Next() {
		todo := Chat_list{}
		err = r1.Scan(&todo.Id, &todo.Owner, &todo.Addr, &todo.Name, &todo.Img, &todo.Time, &todo.No_read)
		if err != nil {
			fmt.Println("64:", err)
			return
		}
		id = todo.Id
		fmt.Println("70id:", id)
	}
	fmt.Println("strconv.Itoa(id):", strconv.Itoa(id))
	if strconv.Itoa(id) == "0" {
		r, err2 := db.Exec("insert into chat_list(owner,addr,name,img,time,no_read)values(?, ?,?,?,?,?)", chat.Owner, chat.Addr, chat.Name, img, strTime, 0)
		r2, err2 := db.Exec("insert into chat_list(owner,addr,name,img,time,no_read)values(?, ?,?,?,?,?)", chat.Addr, chat.Owner, username, userimg, strTime, 0)
		if err2 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println(r2)
		id, err1 := r.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println("insert chat_list succ:", id)
		return
	}
	return
}

//增加消息详情
func (chat *Chat) Insert_chat_content() (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	fmt.Println("93allcontent:", chat.Addr_from, chat.Addr_to, chat.Img_from, chat.Img_to, chat.Content)
	r, err := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", chat.Addr_from, chat.Addr_to, chat.Content, chat.Img_from, chat.Img_to, strTime)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id1, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert chat succ:", id1)

	r1, err := db.Exec("update chat_list set time=? where owner=? and addr =? ", strTime, chat.Addr_from, chat.Addr_to)
	r2, err := db.Exec("update chat_list set time=?,no_read=no_read+1 where owner=? and addr =? ", strTime, chat.Addr_to, chat.Addr_from)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println(r1, r2)

	return
}

//发送借用提醒
func SendBorrow(addr_from string, addr_to string, name_from string, name_to string, content string, img_from string, img_to string) (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	r1, err := db.Query("select * from chat_list where addr=? and owner=?", addr_from, addr_to)
	for r1.Next() {
		todo := Chat_list{}
		err = r1.Scan(&todo.Id, &todo.Owner, &todo.Addr, &todo.Name, &todo.Img, &todo.Time, &todo.No_read)
		if err != nil {
			fmt.Println("140:", err)
			return
		}
		id = todo.Id
		fmt.Println("144id:", id)
	}
	fmt.Println("strconv.Itoa(id):", strconv.Itoa(id))
	if strconv.Itoa(id) == "0" {
		r, err2 := db.Exec("insert into chat_list(owner,addr,name,img,time,no_read)values(?, ?,?,?,?,?)", addr_from, addr_to, name_from, img_to, strTime, 0)
		r2, err2 := db.Exec("insert into chat_list(owner,addr,name,img,time,no_read)values(?, ?,?,?,?,?)", addr_to, addr_from, name_to, img_from, strTime, 0)
		if err2 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println(r2)
		id0, err1 := r.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println("160insert chat_list succ:", id0)

		r, err0 := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", addr_from, addr_to, content, img_from, img_to, strTime)
		if err0 != nil {
			fmt.Println("exec failed, ", err0)
			return
		}
		id1, err1 := r.LastInsertId()
		if err1 != nil {
			fmt.Println("exec failed, ", err1)
			return
		}
		fmt.Println("insert chat succ:", id1)
		r3, err3 := db.Exec("update chat_list set time=? where owner=? and addr =? ", strTime, addr_from, addr_to)
		r4, err4 := db.Exec("update chat_list set time=?,no_read=no_read+1 where owner=? and addr =? ", strTime, addr_to, addr_from)
		if err3 != nil {
			fmt.Println("exec failed, ", err3)
			return
		}
		if err4 != nil {
			fmt.Println("exec failed, ", err4)
			return
		}
		fmt.Println(r3, r4)
		return
	} else {
		r, err1 := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", addr_from, addr_to, content, img_from, img_to, strTime)
		if err1 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		id1, err2 := r.LastInsertId()
		if err2 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println("insert chat succ:", id1)
		r3, err3 := db.Exec("update chat_list set time=? where owner=? and addr =? ", strTime, addr_from, addr_to)
		r4, err3 := db.Exec("update chat_list set time=?,no_read=no_read+1 where owner=? and addr =? ", strTime, addr_to, addr_from)
		if err3 != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		fmt.Println(r3, r4)
		return
	}

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
