package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ColorForm 定义用于绑定复选框的结构体
type ColorForm struct {
	Colors []string `form:"colors[]" binding:"required"` // 绑定复选框数据，必填至少一个选项
}

// handleColors 处理颜色选择请求
func handleColors(c *gin.Context) {
	var form ColorForm
	// 使用 ShouldBind 绑定表单数据
	// 复选框数据以 colors[]=value 格式提交，绑定到 Colors 切片
	if err := c.ShouldBind(&form); err != nil {
		// 如果绑定失败（例如未勾选任何选项），返回400状态码
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 返回绑定后的颜色列表
	c.JSON(200, gin.H{
		"selected_colors": form.Colors,
	})
}

func main() {
	// 创建 Gin 路由器
	r := gin.Default()
	r.Use(cors.Default())
	// 注册路由处理颜色选择请求
	r.POST("/colors", handleColors)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("服务器启动失败: " + err.Error())
	}
}
