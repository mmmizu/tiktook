package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktook/models"
	"tiktook/service"
	"time"
)

type UserLoginResponse struct {
	models.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	models.Response
	User models.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	token := username + password
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{0, "注册成功"},
		UserId:   user.Id,
		Token:    token,
	})
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"})
		return
	}
	token := username + password
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "OK"},
		UserId:   user.Id,
		Token:    token,
	})
}
func UserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "request is invalid",
		})
		return
	}
	user, err := service.UserInfoByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "Ok"},
		User: models.User{
			Id:             user.Id,
			Name:           user.Name,
			Password:       user.Password,
			FollowCount:    10,
			FollowerCount:  10,
			IsFollow:       false,
			TotalFavorited: 10,
			WorkCount:      10,
			FavoriteCount:  10,
			CreatedAt:      time.Now(),
		},
	})
}
