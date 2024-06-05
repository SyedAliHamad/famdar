package routes

import (
	"github.com/famdar/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user/stats/:id", controllers.GetUserStats)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.GET("/users/nearby", controllers.GetUsersNearby)

	return router
}
