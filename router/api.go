package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weet/controller"
)

func apiRouter(api *gin.RouterGroup) {

	// sampleAPI
	api.GET("/sample/:sample_id", controller.GetSampleById)
}
