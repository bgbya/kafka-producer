package c

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"time"
//)
//
//func Hello(c *gin.Context) {
//	c.String(500, "hello")
//	c.JSON(http.StatusOK, gin.H{
//		"name": "tom",
//		"age":  "20",
//	})
//}
//
//type User struct {
//	UserName string `json:"user_name"`
//	PassWD   string `json:"pass_wd"`
//}
//
//func Login(c *gin.Context) {
//	u := User{}
//	err := c.ShouldBind(&u)
//	if err != nil {
//		log.Println("bind error", err)
//	}
//	if u.UserName == "whf" && u.PassWD == "123" {
//		c.JSON(200, gin.H{
//			"status": "success",
//		})
//
//	} else {
//		c.JSON(403, gin.H{
//			"status": "用户名或密码不正确",
//		})
//	}
//}
//func Login2(c *gin.Context) {
//	name := c.Query("name")
//	c.JSON(200, gin.H{
//		"name": name,
//	})
//}
//func MyMiddeleware2(c *gin.Context) {
//	fmt.Println("我的第二个中间件")
//}
//
//func costTime() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		now := time.Now()
//		context.Next()
//		costTime2 := time.Since(now)
//		url := context.Request.URL.String()
//		fmt.Printf("url:%s cost %v", url, costTime2)
//	}
//}
//func main() {
//
//}
//func main2() {
//	e := gin.Default()
//	e.GET("/hello", Hello)
//	e.POST("/login", MyMiddeleware2, Login)
//	e.Use(MyMiddeleware2)
//	e.GET("/login2", Login2)
//	r := gin.Default()
//	admin := r.Group("/admin")
//	admin.Use(gin.BasicAuth(gin.Accounts{
//		"admin": "123456",
//	}))
//	admin.GET("/index", costTime(), func(context *gin.Context) {
//		context.JSON(200, "首页")
//	})
//	go func() {
//		e.Run(":8081")
//	}()
//	r.Run(":8082")
//}
//package main
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"time"
//)
//
//func Hello(c *gin.Context) {
//	c.String(500, "hello")
//	c.JSON(http.StatusOK, gin.H{
//		"name": "tom",
//		"age":  "20",
//	})
//}
//
//type User struct {
//	UserName string `json:"user_name"`
//	PassWD   string `json:"pass_wd"`
//}
//
//func Login(c *gin.Context) {
//	u := User{}
//	err := c.ShouldBind(&u)
//	if err != nil {
//		log.Println("bind error", err)
//	}
//	if u.UserName == "whf" && u.PassWD == "123" {
//		c.JSON(200, gin.H{
//			"status": "success",
//		})
//
//	} else {
//		c.JSON(403, gin.H{
//			"status": "用户名或密码不正确",
//		})
//	}
//}
//func Login2(c *gin.Context) {
//	name := c.Query("name")
//	c.JSON(200, gin.H{
//		"name": name,
//	})
//}
//func MyMiddeleware2(c *gin.Context) {
//	fmt.Println("我的第二个中间件")
//}
//
//func costTime() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		now := time.Now()
//		context.Next()
//		costTime2 := time.Since(now)
//		url := context.Request.URL.String()
//		fmt.Printf("url:%s cost %v", url, costTime2)
//	}
//}
//func main() {
//
//}
//func main2() {
//	e := gin.Default()
//	e.GET("/hello", Hello)
//	e.POST("/login", MyMiddeleware2, Login)
//	e.Use(MyMiddeleware2)
//	e.GET("/login2", Login2)
//	r := gin.Default()
//	admin := r.Group("/admin")
//	admin.Use(gin.BasicAuth(gin.Accounts{
//		"admin": "123456",
//	}))
//	admin.GET("/index", costTime(), func(context *gin.Context) {
//		context.JSON(200, "首页")
//	})
//	go func() {
//		e.Run(":8081")
//	}()
//	r.Run(":8082")
//}
