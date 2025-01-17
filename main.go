package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TODO struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

var todos []TODO

func main() {
	r := gin.Default()
	//添加 TODO
	r.POST("/todo", func(c *gin.Context) {
		var todo TODO
		c.BindJSON(&todo)
		todos = append(todos, todo)
		fmt.Println(todos)
		c.JSON(200, gin.H{"status": "ok"})
	})
	//修改 TODO
	r.PUT("/todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		var todo TODO
		c.BindJSON(&todo)
		todos[index] = todo
		c.JSON(200, gin.H{"status": "ok"})
	})
	//获取TODO
	r.GET("/todo", func(c *gin.Context) {
		c.JSON(200, todos)
	})
	//查询 TODO
	r.GET("/todo/：index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		c.JSON(200, todos[index])
	})
	r.Run(":8080")

}
