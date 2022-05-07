package db

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func Init() (err error) {
	// 打开mysql数据库
	fmt.Println()
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sharefish?charset=utf8&parseTime=true")

	if err != nil {
		panic(err)
	}
	//err = db.Ping()
	if err != nil {
		panic(err)
		fmt.Print("数据库链截失败")
	}
	fmt.Println("已经连接MYSQL")
	return
}

//查看需求
func ListNeedAll() (DemandList []DemandDB, err error) {
	// 查询数据
	var sqlStr = `select * from demand `
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("select_demand_err:", err)
		return nil, err
	}

	for rows.Next() {
		DemandStr := DemandDB{}
		err1 := rows.Scan(&DemandStr.DemandID, &DemandStr.DemandKinds, &DemandStr.DemandName, &DemandStr.DemandTime, &DemandStr.DemandAddr)

		if err1 != nil {
			fmt.Println("select_demand_scan_err:", err1)
			return nil, err
		}
		DemandList = append(DemandList, DemandStr)
	}
	return
}

//通过id查看需求
func IdListNeeds(id int64) (DemandList []DemandDB, err error) {
	var sqlStr = `select * from demand where demandID =?`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		DemandStr := DemandDB{}
		err = rows.Scan(&DemandStr.DemandID, &DemandStr.DemandKinds, &DemandStr.DemandName, &DemandStr.DemandTime, &DemandStr.DemandAddr)

		if err != nil {
			return nil, err
		}
		DemandList = append(DemandList, DemandStr)
	}

	return
}

//查看我的需求
func MyListNeeds(demandAddr common.Address) (DemandList []DemandDB, err error) {
	// 查询数据
	demandAddrStr := demandAddr.Hex()

	var sqlStr = `select * from demand where demandAddr =?`
	rows, err := db.Query(sqlStr, demandAddrStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		DemandStr := DemandDB{}
		err = rows.Scan(&DemandStr.DemandID, &DemandStr.DemandKinds, &DemandStr.DemandName, &DemandStr.DemandTime, &DemandStr.DemandAddr)

		if err != nil {
			return nil, err
		}
		DemandList = append(DemandList, DemandStr)
	}

	return
}

//需求总数
func MyCountNeeds(demandAddr common.Address) (DemandList []DemandDB, err error) {
	// 查询数据
	demandAddrStr := demandAddr.Hex()

	var sqlStr = `select count(demandID) from demand where demandAddr =?`
	rows, err := db.Query(sqlStr, demandAddrStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		DemandStr := DemandDB{}
		err = rows.Scan(&DemandStr.CountDemand)

		if err != nil {
			return nil, err
		}
		DemandList = append(DemandList, DemandStr)
	}

	return
}

//插入需求数据
func InsertDemand(DemandKind string, DemandAddr string, DemandName string) (err error) {
	nowTime := time.Now()
	var sqlStr = `INSERT INTO demand(demandKinds,demandName,demandTime,demandAddr) VALUES (?,?,?,?)`
	_, err = db.Exec(sqlStr, DemandKind, DemandName, nowTime, DemandAddr)

	if err != nil {
		fmt.Println("连接失败", err)
		return err
	}
	fmt.Println("插入数据成功")
	return nil
}

//删除需求数据
func DeleteNeeds(id int64) (DemandList []DemandDB, err error) {
	var sqlStr = `delete from demand where demandID= ?`
	_, err = db.Exec(sqlStr, id)
	return
}

//发送同意借用消息
func AgreeBorrow(addr_from string, addr_to string, content string, img_from string, img_to string) (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	r, err1 := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", addr_from, addr_to, content, img_from, img_to, strTime)
	if err1 != nil {
		fmt.Println("exec failed, ", err1)
		return
	}
	id1, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}
	fmt.Println("insert chat succ:", id1)
	r3, err3 := db.Exec("update chat_list set time=? where owner=? and addr =? ", strTime, addr_from, addr_to)
	r4, err3 := db.Exec("update chat_list set time=?,no_read=no_read+1 where owner=? and addr =? ", strTime, addr_to, addr_from)
	if err3 != nil {
		fmt.Println("exec failed, ", err3)
		return
	}
	fmt.Println(r3, r4)
	return
}

//发送不同意借用消息
func DisagreeBorrow(addr_from string, addr_to string, content string, img_from string, img_to string) (err error) {
	times := time.Now()
	strTime := times.Format("15:05:05")
	r, err1 := db.Exec("insert into chat(addr_from,addr_to,content,img_from,img_to,time)values(?, ?,?,?,?,?)", addr_from, addr_to, content, img_from, img_to, strTime)
	if err1 != nil {
		fmt.Println("exec failed, ", err1)
		return
	}
	id1, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Println("exec failed, ", err2)
		return
	}
	fmt.Println("insert chat succ:", id1)
	r3, err3 := db.Exec("update chat_list set time=? where owner=? and addr =? ", strTime, addr_from, addr_to)
	r4, err3 := db.Exec("update chat_list set time=?,no_read=no_read+1 where owner=? and addr =? ", strTime, addr_to, addr_from)
	if err3 != nil {
		fmt.Println("exec failed, ", err3)
		return
	}
	fmt.Println(r3, r4)
	return
}
