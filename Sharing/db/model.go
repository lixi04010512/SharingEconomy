package db

import "time"

type DemandDB struct {
	DemandID    int       `sql:"demandID" json:"demandID from:"demandID'`
	DemandKinds string    `sql:"demandKinds" json:"demandKinds" from:"demandKinds"`
	DemandName  string    `sql:"demandName" json:"demandName" from:"demandName"`
	DemandTime  time.Time `sql:"demandTime" json:"demandTime" from:"demandTime"`
	DemandAddr  string    `sql:"demandAddr" json:"demandAddr" from:"demandAddr"`
}
