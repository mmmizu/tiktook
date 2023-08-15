package models

import "time"

type User struct {
	Id             int64     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Password       string    `json:"password,omitempty"`
	FollowCount    int64     `json:"follow_count,omitempty"`
	FollowerCount  int64     `json:"follower_count,omitempty"`
	IsFollow       bool      `json:"is_follow,omitempty" gorm:"-"`
	TotalFavorited int64     `json:"total_favorited,omitempty"`
	WorkCount      int64     `json:"work_count,omitempty"`
	FavoriteCount  int64     `json:"favorite_count,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
