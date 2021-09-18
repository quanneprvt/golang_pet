package Game

import "go_backend/src/services/game"
import "github.com/gin-gonic/gin"
import "net/http"

func GetGame(c *gin.Context) {
	var albums = Game.Get()
	c.IndentedJSON(http.StatusOK, albums)
}
