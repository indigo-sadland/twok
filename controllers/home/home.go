package home

import (
	"github.com/gin-gonic/gin"
	"twok/core/router"
)

func Load() {
	r := router.GetRouter()
	r.GET("/", HomeGet)
}

func HomeGet(c *gin.Context) {
	c.HTML(200, "home.html", "")
}
