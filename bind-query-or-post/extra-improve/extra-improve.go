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
