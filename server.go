package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/saimohanreddyt/golang-gin-poc/service"
)
var (
	videoServices service.videoService = server.New()
	videoController controller.VideoController = controller.New(videoServices)
)
func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}