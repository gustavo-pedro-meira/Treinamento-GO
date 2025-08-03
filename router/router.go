package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	v1.GET("/opening", func(c *gin.Context){
		c.JSON(http.StatusOK, "GET Opening")
	})

	v1.POST("/opening", func(c *gin.Context){
		c.JSON(http.StatusCreated, "POST Opening")
	})

	v1.PUT("/opening", func(c *gin.Context){
		c.JSON(http.StatusOK, "PUT Opening")
	})

	v1.DELETE("/opening", func(c *gin.Context){
		c.JSON(http.StatusOK, "DELETE Opening")
	})

	v1.GET("/openings", func(c *gin.Context){
		c.JSON(http.StatusOK, "GET Openings")
	})
}