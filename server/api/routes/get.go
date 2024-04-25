package routes

import (
	"strconv"

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

func getUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		// Get the id
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid ID",
			})
		}

		// Check id is above 0 for uint conversion
		if id > 0 {
			db.Model(models.User{UserID: uint(id)}).First(&user)
			c.JSON(200, user)
		}

		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
	}
}

// Loads the endpoints for all "GET" requests
func LoadGetEndpoints(router *gin.Engine, database *gorm.DB) {
	endpointGroup := router.Group("/api/get")
	{
		endpointGroup.GET("/users", getAllUsers(database))
		endpointGroup.GET("/user/:id", getUserByID(database))
	}
}
