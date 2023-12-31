package main

import (
	"github.com/gin-gonic/gin"
	"tiktook/models"
)

func main() {
	r := gin.Default()
	initRouter(r)
	models.DbInit()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
