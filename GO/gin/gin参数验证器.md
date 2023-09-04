required 必填字段  binding:"required"
min 最小长度   min=3
max 最长长度  max=10
len 长度   len=6

eq
ne
gt
get
lt
lte;
eqfield  等于其他字段的值
nefield  不等于其他字段的值
- 忽略字段  binding:"-"


```go

package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name       string `json:"name" binding:"min=5"`
	Age        int    `json:"age" binding:"gt=18,lt=35"`
	Password   string `json:"password"`
	Repassword string `json:"repassword" binding:"eqfield=Password"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.String(203, err.Error())
			return
		}
		ctx.JSON(200, gin.H{"data": user})
	})

	r.Run(":8080")
}




```

// 枚举
oneof=red green

contains=ck  包含ck字符串
excludes=   //不包含
startswith   //字符串前缀
endswiith    //字符串后缀


dive  // 数组

// 网络验证
ip
ipv4
ipv6
uri
url

//日期验证
datetime=2006-01-02 15:04:05





自定义错误消息
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type User struct {
	Name string `json:"name" binding:"required" msg:"用户校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

func test(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {

		c.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
		return

	}
	c.JSON(200, gin.H{"data": user})

}

func main() {
	r := gin.Default()
	r.POST("/", test)
	r.Run(":8080")
}

```