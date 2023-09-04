```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"` //form query
	Age  int    `json:"age" form:"age" uri:"age"`    // url shoulduri
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	r := gin.Default()

	// 绑定json
	r.POST("/", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, user)
	})
	// query param
	r.POST("/query", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindQuery(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, user)
	})
	// 绑定uri 127.0.0.1:8080/uri/ck/21/男
	r.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, user)
	})
	// form-data / x-www-form-urlencode
	// 默认的tag 就是form
	r.POST("/form", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}

```