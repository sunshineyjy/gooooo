package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	//有一个方法可以匹配 /user/tom, 也可以匹配 /user/tom/send
	//如果没有任何了路由匹配 /user/tom, 它将会跳转到 /user/tom/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "guest")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/person_add", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	// Example for binding JSON ({"username": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json User

		if err := c.ShouldBindJSON(&json); err == nil {
			if json.Username == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "username": json.Username, "pass": json.Password})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Example for binding a HTML form (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var form User
		if err := c.ShouldBind(&form); err == nil {
			if form.Username == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in 2"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized 2"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// r.POST("/upload", func(c *gin.Context) {
	// 	// Single file
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	c.SaveUploadedFile(file, dst)

	// 	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// })

	// r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// r.POST("/upload", func(c *gin.Context) {
	// 	// Multipart form
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["upload[]"]

	// 	for _, file := range files {
	// 		log.Println(file.Filename)

	// 		// Upload the file to specific dst.
	// 		c.SaveUploadedFile(file, dst)
	// 	}
	// 	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	// })

	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8088")
	r.Run(":8088")
}
