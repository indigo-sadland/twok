package home

import (
	"github.com/gin-gonic/gin"
	"github.com/indigo-sadland/twok/core/router"
)

func Load() {
	r := router.GetRouter()
	r.GET("/", Get)
}

func Get(c *gin.Context) {
	c.HTML(200, "home.html", "")
}
