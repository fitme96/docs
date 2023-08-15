```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _getlist(c *gin.Context) {
	articlelist := []ArticleModel{
		{"go", "this go"},
		{"python", "this python"},
		{"java", "this java"},
	}
	c.JSON(200, Response{0, articlelist, "成功"})
}

// 获取详情
func _getdetail(c *gin.Context) {
	// 获取param中的id
	fmt.Println(c.Param("id"))
	article := []ArticleModel{
		{"go", "this go"},
	}
	c.JSON(200, Response{2, article, "成功"})
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contextType := c.GetHeader("Content-Type")
	switch contextType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 创建文章

func _create(c *gin.Context) {
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, Response{0, article, "添加成功"})
}
func _update(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func _delete(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func main() {
	r := gin.Default()
	r.GET("/articles", _getlist)
	r.GET("/articles/:id", _getdetail)
	r.POST("/articles", _create)
	r.PUT("/articles/:id", _update)
	r.DELETE("/aarticles/:id", _delete)
	r.Run(":8080")
}

```