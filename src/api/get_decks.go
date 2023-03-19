package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDecksHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.cardService.GetDecks())
	}
}
