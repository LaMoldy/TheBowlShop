package routes

import "github.com/gin-gonic/gin"

func testEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func LoadGetEndpoints(router *gin.Engine) {
	endpointGroup := router.Group("/api")
	{
		endpointGroup.GET("/test", testEndpoint)
	}
}
