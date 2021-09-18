package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	GameController "go_backend/src/controllers/game"
	"net/http"
)

func InitApp() {
	fmt.Println("Hello from app")
	router := gin.Default()
	//API Define
	router.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Health check OK")
	})
	router.GET("/game", GameController.GetGame)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run("localhost:8080")
}
