package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		// c.String(http.StatusOK, "hello world")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  "hello world",
		})
	})

	v1 := r.Group("v1")

	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  param,
		})
	})

	v1.POST("/user", func(ctx *gin.Context) {
		var data Test

		ctx.BindJSON(&data)

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  data,
		})
	})

	v1.GET("/user", func(ctx *gin.Context) {
		query := ctx.Query("name")
		// c.String(http.StatusOK, "hello world")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  query,
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run()
}
