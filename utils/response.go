package utils

import "github.com/gin-gonic/gin"

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
