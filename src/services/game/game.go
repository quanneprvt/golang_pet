package Game

import "go_backend/src/models"
import "github.com/gin-gonic/gin"

type Game Models.Game

func Init(ct *gin.Context, cn string) {
	var c Game
	c.Description = ""
	c.Played = 0
	c.Title = ""
}

func Get() Game {
	return Game{"123", "564", "1454", 0}
}
