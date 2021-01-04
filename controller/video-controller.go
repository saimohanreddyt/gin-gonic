package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/saimohanreddyt/golang-gin-poc/entity"
	)

type VideoController interface {
	Findall() []entity.Video
	Save(ctx *gin.Context) entity.Video

}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller {
		service: service,
	}
}

func (c *controller) Findall() []entity.Video {
	return service().FindAll()

}
func (c *controller) Save(ctx *gin.Context) entity.Video{
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}