package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeckRequestBody struct {
	Shuffled bool   `json:"shuffled"`
	Cards    string `json:"cards"`
}

func PostDeckHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data DeckRequestBody
		c.Bind(&data)

		deck := handler.cardService.CreateNewDeck(data.Shuffled, data.Cards)

		c.JSON(http.StatusOK, deck)
	}
}
