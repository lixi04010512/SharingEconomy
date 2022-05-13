package db

import (
	"time"
)

type DemandDB struct {
	DemandID    int64     `sql:"demandID" json:"demandID"     from:"demandID"`
	DemandKinds string    `sql:"demandKinds" json:"demandKinds" from:"demandKinds"`
	DemandName  string    `sql:"demandName" json:"demandName" from:"demandName"`
	DemandTime  time.Time `sql:"demandTime" json:"demandTime" from:"demandTime"`
	DemandAddr  string    `sql:"demandAddr" json:"demandAddr" from:"demandAddr"`
	CountDemand int64     `sql:"count(demandID)" json:"countDemand" from:"countDemand"`
}

type Chat_list struct {
	Id      int    `db:"id" json:"id" form:"id"`
	Owner   string `db:"owner" json:"owner" form:"owner"`
	Addr    string `db:"addr" json:"addr" form:"addr"`
	Name    string `db:"name" json:"name" form:"name"`
	Img     string `db:"img" json:"img" form:"img"`
	Time    string `db:"time" json:"time" form:"time"`
	No_read int    `db:"no_read" json:"no_read" form:"no_read"`
}
type Chat struct {
	Id        int    `db:"id" json:"id" form:"id"`
	Addr_from string `db:"addr_from" json:"addr_from" form:"addr_from"`
	Addr_to   string `db:"addr_to" json:"addr_to" form:"addr_to"`
	Content   string `sb:"content" json:"content" form:"content"`
	Img_from  string `db:"img_from" json:"img_from" form:"img_from"`
	Img_to    string `db:"img_to" json:"img_to" form:"img_to"`
	Time      string `db:"time" json:"time" form:"time"`
	Confirm   string `db:"confirm" json:"confirm" form:"confirm"`
}
