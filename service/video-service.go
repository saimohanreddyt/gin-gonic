package service

import "gitlab.com/saimohanreddyt/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FinAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

type New() VideoService{
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FinAll() []entity.Video {
	return services.videos

}
