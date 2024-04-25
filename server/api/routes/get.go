package routes

import (
	"github.com/LaMoldy/TheBowlShop/internal/database/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		db.First(&users)
		c.JSON(200, users)
	}
}

func getUserByEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		email := c.Param("email")
		db.Model(models.User{Email: email}).First(&user)
		c.JSON(200, user)
	}
}

func LoadGetEndpoints(router *gin.Engine, database *gorm.DB) {
	endpointGroup := router.Group("/api/get")
	{
		endpointGroup.GET("/users", getAllUsers(database))
		endpointGroup.GET("/user/:email", getUserByEmail(database))
	}
}
