package main

import (
	_ "dbmanager/sqldb"

	"github.com/go-phd/ssf"
)

/*
type Cdbmanager struct {
	//ssf.Cssf
	stopping bool
}
*/
//var dm Cdbmanager

func Init() {
	ssf.Logger.Infoln("init success.")
}
