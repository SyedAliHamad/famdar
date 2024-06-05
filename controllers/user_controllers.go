package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/famdar/database"
	"github.com/famdar/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserStats retrieves stats for a specific user by ID
func GetUserStats(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// You can add any additional stats computation here if needed
	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
	})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUsersNearby fetches users within a specified radius from a given location
func GetUsersNearby(c *gin.Context) {
	// Parse query parameters
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}
	lng, err := strconv.ParseFloat(c.Query("lng"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}
	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
		return
	}

	log.Println(lat,lng, radius)
	// Query users within the specified radius using PostGIS
	var nearbyUsers []models.User
	err = database.DB.Raw(`
		SELECT *
		FROM users
		WHERE ST_Distance(location, ST_MakePoint(?, ?)::geography) <= ?
	`, lng, lat, radius).Find(&nearbyUsers).Error
	log.Println(nearbyUsers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby users"})
		return
	}

	c.JSON(http.StatusOK, nearbyUsers)
}
