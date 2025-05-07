package main

import (
    "io"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/gin-contrib/cors"
)

// 定义一个结构体用于绑定 JSON 数据
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // 创建 Gin 路由
    r := gin.Default()

	r.Use(cors.Default())

    // 路由 1: 使用 c.Request.Body 读取原始请求体
    r.POST("/raw-body", func(c *gin.Context) {
        // 读取 c.Request.Body 的内容
        bodyBytes, err := io.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(400, gin.H{"error": "Failed to read body"})
            return
        }
        // 将读取到的原始数据作为字符串返回
        c.JSON(200, gin.H{"raw_body": string(bodyBytes)})
    })

    // 路由 2: 使用 c.ShouldBindBodyWith 绑定 JSON 到结构体
    r.POST("/bind-body", func(c *gin.Context) {
        var user User
        // 使用 JSON 绑定器将请求体绑定到 user 结构体
        if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
            c.JSON(400, gin.H{"error": "Failed to bind JSON"})
            return
        }
        // 返回绑定后的结构体数据
        c.JSON(200, gin.H{"name": user.Name, "email": user.Email})
    })

    // 启动服务器，监听 8080 端口
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
