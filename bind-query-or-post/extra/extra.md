# 关于这个额外的示例

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")

```



运行代码后，使用命令测试：

```bash
# bind by query
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz"
# bind by json
curl -X GET localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"
```

gin框架分别打印了如下信息：

![](/home/Moranis/.config/marktext/images/2025-05-12-20-04-52-image.png)

![](/home/Moranis/.config/marktext/images/2025-05-12-20-05-40-image.png)

gin对不同请求的数据类型做出了不同的回应，原因是：

- **c.Bind(&person)**：
  - c.Bind是一个通用绑定方法，会根据请求的上下文（例如请求方法和数据格式）自动选择合适的绑定器。
  - 对于GET请求，c.Bind默认尝试绑定查询字符串（query）或表单数据（form-data），因为GET请求通常不携带请求体。
  - 它依赖于结构体字段的form标签（例如form:"name"）来映射参数。
- **c.BindJSON(&person)**：
  - c.BindJSON专门用于绑定JSON格式的请求体数据。
  - 它会检查Content-Type头信息（例如application/json）来确定是否处理JSON数据。
  - 依赖于结构体字段的json标签（例如json:"name"）来映射JSON字段。

当然，由于**c.Bind**也可以绑定json,这段代码应该优先判断数据类型，可以将处理函数修改为这样：

```go
func startPage(c *gin.Context) {
	var person Person

	// 检查Content-Type是否为JSON
	if c.GetHeader("Content-Type") == "application/json" {
		if err := c.BindJSON(&person); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON: " + err.Error()})
			return
		}
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	} else {
		if err := c.Bind(&person); err != nil {
			c.JSON(400, gin.H{"error": "Invalid query string: " + err.Error()})
			return
		}
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}
```

这样代码更有逻辑。
