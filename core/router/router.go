package router

import (
	"github.com/gin-gonic/gin"
)

var r Router

type Router struct {
	Do *gin.Engine
}

func SetRouter() Router {
	// Initialize gin router and serve frontend static files
	router := gin.Default()
	r = Router{
		Do: router,
	}

	// Setup route group for API
	//api := router.Group()

	return r
}

func GetRouter() *gin.Engine {
	return r.Do
}
