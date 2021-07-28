package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//构建模型
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

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
	db.AutoMigrate(&User{}) //自动创建数据表  相当于将模型与数据库进行自动绑定   则运行时会自动在数据库中创建相应字段的表  且表的名字为结构体（模型）的小写复数
	return db
}

//初始化数据库方法2
func Init() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8080)/ginessential?charsetTime=true")
	db.AutoMigrate(&User{}) //自动创建数据表  相当于将模型与数据库进行自动绑定   则运行时会自动在数据库中创建相应字段的表  且表的名字为结构体（模型）的小写复数
	return db
}

//判断电话号码是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//随机生成一个长度为10的字符串
func RandomString(n int) string {
	var letters = []byte("hdfjigduysvhdfsbhjdsfgyudguavbajxiofheruibcfd")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix()) //随机数种子
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//主函数
func main() {
	//初始化数据库
	db := InitDB()
	defer db.Close()

	//创建路由和请求
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//如果名字没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)

		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})

	//运行服务器
	panic(r.Run(":8080"))

}
