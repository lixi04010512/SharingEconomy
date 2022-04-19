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
	db, err = sql.Open("mysql", "root:ys3285739@tcp(127.0.0.1:3306)/sharefish?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		fmt.Print("数据库链截失败")
	}
	fmt.Println("已经连接MYSQL")
	var sqlStr = `DROP TABLE IF EXISTS demand;
CREATE TABLE demand  (
  demandID int(0) NOT NULL AUTO_INCREMENT,
  demandKinds varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  demandName longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  demandTime timestamp(0) NULL DEFAULT NULL,
  demandAddr varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (demandID) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
`
	_, err = db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return
}

//查看需求
func ListNeeds() (DemandList []DemandDB, err error) {
	// 查询数据\
	var sqlStr = `SELECT * FROM demand `
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		DemandStr := DemandDB{}
		err = rows.Scan(&DemandStr.DemandID, &DemandStr.DemandKinds, &DemandStr.DemandAddr, &DemandStr.DemandName, &DemandStr.DemandTime)
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
		fmt.Println(&DemandStr.DemandID)
		if err != nil {
			return nil, err
		}
		DemandList = append(DemandList, DemandStr)
	}

	return
}

//编辑需求数据
func editDemand(DemandID int, DemandKind string, DemandName string) (err error) {
	var sqlStr = `updata demand set demandKinds = ` + DemandKind + `,demandName =` + DemandName + ` where demandID = ?`
	_, err = db.Query(sqlStr, DemandID)
	return err
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
func deleteNeeds(id int) (err error) {
	var sqlStr = `delete from demand where demandID= ?`
	_, err = db.Exec(sqlStr, id)
	return err
}