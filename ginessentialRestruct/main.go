package main

import (
	"ProjectPractice/ginessentialRestruct/common"
	"ProjectPractice/ginessentialRestruct/control"
	"ProjectPractice/ginessentialRestruct/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//主函数
func main() {
	//初始化数据库
	db := common.InitDB()
	defer db.Close()

	//创建路由和请求
	r := gin.Default()
	r.POST("/register", control.Register)
	r.POST("/login", control.Login)
	r.GET("/info", middleware.AuthMiddleware(), control.Info)
	r.GET("hot-fix-test", control.HotfixTest)

	//运行服务器
	panic(r.Run(":8080"))
	fmt.Print("master test...")
}
