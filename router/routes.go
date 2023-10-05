package router

import (
	handler "github.com/NatanColleoni/gop/handler/opening"
	gin "github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	var v1 *gin.RouterGroup = router.Group("/api/v1")

	{
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
