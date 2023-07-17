package routes

import "github.com/gin-gonic/gin"

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/", func(ctx *gin.Context) {
		isValidated := false

		if isValidated {
			ctx.AbortWithStatusJSON(400, gin.H{
				"massage": "bad requrest, some field no value",
			})

			return
		}

		ctx.JSON(200, gin.H{
			"hello": "world",
		})

		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})
}
