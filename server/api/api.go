package api

import (
	"github.com/LaMoldy/TheBowlShop/api/routes"
	"github.com/gin-gonic/gin"
)

// Creates an api
func CreateApi() (*gin.Engine, error) {
	// Create router
	router := gin.New()

	// Add the middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Set trusted proxies
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Load the routes
	routes.LoadGetEndpoints(router)

	return router, nil
}
