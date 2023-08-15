package controller

import (
	"github.com/gin-gonic/gin"
	"tiktook/models"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func Publish(c *gin.Context)     {}
func PublishList(c *gin.Context) {}
