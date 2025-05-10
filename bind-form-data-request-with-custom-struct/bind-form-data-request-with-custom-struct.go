package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StructA 定义一个简单的结构体
type StructA struct {
	FieldA string `form:"field_a" binding:"required"` // 映射到表单的 field_a，必填
}

// StructB 包含嵌套的 StructA 和一个表单字段
type StructB struct {
	NestedStruct StructA // 嵌套结构体，无 form 标签
	FieldB       string  `form:"field_b" binding:"required"` // 映射到表单的 field_b，必填
}

// handleGetDataB 处理 /getb 路由
func handleGetDataB(c *gin.Context) {
	var b StructB
	// 使用 ShouldBind 绑定查询参数或表单数据
	if err := c.ShouldBind(&b); err != nil {
		// 如果绑定失败（例如缺少必填字段），返回400状态码
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 返回绑定后的数据
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func main() {
	// 创建 Gin 路由器
	r := gin.Default()

	// 添加CORS中间件，允许跨域请求
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000", "http://127.0.0.1:8000"}, // 允许多个来源
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},                         // 允许的HTTP方法，包括OPTIONS（预检请求）
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},               // 允许的请求头
		AllowCredentials: true,                                                       // 允许携带凭据
	}))

	// 注册路由
	r.GET("/getb", handleGetDataB)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("服务器启动失败: " + err.Error())
	}
}
