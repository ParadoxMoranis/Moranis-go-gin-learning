package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Person 定义用于绑定查询字符串的结构体
type Person struct {
	Name     string    `form:"name" binding:"required"`                              // 姓名，必填
	Email    string    `form:"email" binding:"required,email"`                       // 邮箱，必填且需符合邮箱格式
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" binding:"required"` // 生日，必填，格式为YYYY-MM-DD
}

func main() {
	// 创建默认Gin路由器，包含日志和恢复中间件
	route := gin.Default()

	// 注册GET路由，处理/api/person请求
	route.GET("/api/person", handlePerson)

	// 启动服务器，监听8085端口
	if err := route.Run(":8085"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handlePerson 处理GET请求，绑定查询字符串并返回结果
func handlePerson(c *gin.Context) {
	var person Person

	// 尝试绑定查询字符串到Person结构体
	// c.ShouldBind处理GET请求的查询字符串，自动映射form标签对应的参数
	// binding标签用于验证：required表示必填，email验证邮箱格式
	if err := c.ShouldBind(&person); err != nil {
		// 绑定失败（例如缺少必填字段或格式错误），返回400错误和错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 绑定成功，返回200状态码和绑定的数据
	c.JSON(http.StatusOK, gin.H{
		"message":  "Data received successfully",
		"name":     person.Name,
		"email":    person.Email,
		"birthday": person.Birthday.Format("2006-01-02"),
	})
}
