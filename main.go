package main

import (
	"github.com/gin-gonic/gin"
	"todoapp/server"
)

func main() {
	router := gin.Default()

	router.GET("/", server.GetList)
	router.POST("/", server.AddItem)
	router.POST("/select/:id", server.SelectItem)
	router.POST("/done/:id", server.DoneItem)
	router.POST("/remove/:id", server.RemoveItem)
	router.POST("/remove", server.RemoveAll)
	router.POST("/remove-selected", server.RemoveSelecteds)
	router.POST("/done-selected", server.DoneSelecteds)

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static/")

	router.Run("localhost:8080")
}
