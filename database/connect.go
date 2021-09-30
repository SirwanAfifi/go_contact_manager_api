package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USERNAME = ""
const DB_PASSWORD = ""
const DB_NAME = ""
const DB_HOST = ""
const DB_PORT = ""

var Db *gorm.DB
func InitDb() *gorm.DB {
   Db = connectDB()
   return Db
}

func connectDB() (*gorm.DB) {
	var err error
    dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}