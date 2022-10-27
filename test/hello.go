package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/didip/tollbooth"

	"github.com/didip/tollbooth/limiter"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	fmt.Println("hello")
	//createHttpServer()
	//Default返回一个默认的路由引擎
	r := gin.Default()
	// 使用全局CORS中间件。
	// router.Use(Cors())
	//rate-limit 中间件
	//lmt := tollbooth.NewLimiter(1, nil)
	//lmt.SetMessage("服务繁忙，请稍后再试...")

	//设定请求静态资源路径
	r.StaticFS("/public", http.Dir("C:/Users/laizx02/Desktop/GO/staticResorces"))
	//GET请求不带参数 http(s)://ip:port/ping
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//GET 路径参数请求 例如： http(s)://ip:port/user/Tom
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})
	//GET ?后带参数请求 例如： http(s)://ip:port/users?name=Tom&role=student
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	//POST form表单提交请求 http(s)://ip:port/form
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// GET和POST form表单混合  http://ip:port/posts?id=9876&page=7
	r.POST("/posts", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password", "000000")
		ctx.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	//POST json提交请求|gin方式接收 http://ip:port/postjson
	r.POST("postjson1", func(c *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		err := c.BindJSON(&json)
		if nil != err {
			log.Printf("err=%s",err)
		}
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"username": json["username"],
			"password": json["password"],
		})
	})

	//POST json提交请求|struct方式接收
	r.POST("postjson2", func(ctx *gin.Context) {
		user := User{}
		err := ctx.BindJSON(&user)
		if nil != err {
			log.Printf("err=%s",err)
		}
		username := user.Username
		password := user.Password
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//Map参数(字典参数) curl -g "http://ip:port/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	//重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/goindex")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	//分组路由
	postjson3Handler := func(c *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		err := c.BindJSON(&json)
		if nil != err {
			log.Printf("err=%s",err)
		}
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"username": json["username"],
			"password": json["password"],
		})
	}
	postjson4Handler := func(ctx *gin.Context) {
		user := User{}
		err := ctx.BindJSON(&user)
		if nil != err {
			log.Printf("err=%s",err)
		}
		username := user.Username
		password := user.Password
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	}
	v1 := r.Group("/v1")
	{
		v1.POST("/postjson3", postjson3Handler)
		v1.POST("/postjson4", postjson4Handler)

		// 下面是群组中间的用法
		// v1.Use(Cors())
		// 单个中间件的用法
		// v1.GET("/user/:id/*action",Cors(), api.GetUser)
		// rate-limit
		// v1.GET("/user/:id/*action", LimitHandler(lmt), api.GetUser)
		// v1.GET("/user/:id/*action", Cors(), api.GetUser)
		// AJAX OPTIONS ，下面是有关OPTIONS用法的示例
		// v1.OPTIONS("/users", OptionsUser)      // POST
		// v1.OPTIONS("/users/:id", OptionsUser)  // PUT, DELETE
	}

	//gin html模板实现
	r.LoadHTMLGlob("templates/*")

	user1 := &User{Username: "Geektutu", Password: "123456"}
	user2 := &User{Username: "Jack", Password: "654321"}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.html", gin.H{
			"title":   "Gin",
			"userArr": [2]*User{user1, user2},
		})
	})

	err := r.Run(":8001")
	if err != nil {
		log.Printf("启动gin8001端口失败，err=%s",err)
	}
}
