package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	User="root"
	Password="root"
	Dbname="test"
	Port="3306"
	Host="127.0.0.1"

)
var DbInit=DbWorkerConnect{}

type  DbWorkerConnect struct {
	Conn *gorm.DB
}

func NewWorkerConnect() {
	dbpath :=User+":"+Password+"@("+Host+":"+Port+")/"+Dbname
	db, err := gorm.Open("mysql", dbpath+"?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err!=nil {
		panic(err)
	}
	DbInit.Conn=db
}