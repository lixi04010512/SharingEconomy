package db

import "time"

type Demand struct {
	DemandID   int       `sql:"demandID" json:"demandID from:"demandID'`
	DemandName string    `sql:"demandName" json:"demandName" from:"demandName"`
	DemandKind string    `sql:"demandKind" json:"demandKind" from:"demandKind"`
	DemandTime time.Time `sql:"demandTime" json:"demandTime" from:"demandTime"`
	DemandAddr string    `sql:"demandAddr" json:"demandAddr" from:"demandAddr"`
}
