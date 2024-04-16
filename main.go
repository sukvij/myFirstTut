package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "welcome to the users list......")
	})
	app.Run(":8080")
}
