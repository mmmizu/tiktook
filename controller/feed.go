package controller

import (
	"github.com/gin-gonic/gin"
	"tiktook/models"
)

type FeedResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  models.Response{StatusCode: 0},
	//	VideoList:
	//	NextTime:  time.Now().Unix(),
	//})
}
