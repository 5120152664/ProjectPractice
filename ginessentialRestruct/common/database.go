package common

import (
	"ProjectPractice/ginessentialRestruct/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//初始化数据库方法1
func InitDB() *gorm.DB {
	dirverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(dirverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//db,_:=gorm.Open("mysql","root:root@tcp(127.0.0.1:8080)/ginessential?charsetTime=true")
	db.AutoMigrate(&model.User{}) //自动创建数据表  相当于将模型与数据库进行自动绑定   则运行时会自动在数据库中创建相应字段的表  且表的名字为结构体（模型）的小写复数
	DB = db
	return db
}

//初始化数据库方法2
func Init() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8080)/ginessential?charsetTime=true")
	db.AutoMigrate(&model.User{}) //自动创建数据表  相当于将模型与数据库进行自动绑定   则运行时会自动在数据库中创建相应字段的表  且表的名字为结构体（模型）的小写复数
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
