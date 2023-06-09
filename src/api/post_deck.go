package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewDeckResponse struct {
	Id        string `json:"id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type DeckRequestBody struct {
	Shuffled bool   `json:"shuffled"`
	Cards    string `json:"cards"`
}

func PostDeckHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data DeckRequestBody
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		deck := handler.cardService.CreateNewDeck(data.Shuffled, data.Cards)

		c.JSON(http.StatusOK, NewDeckResponse{
			Id:        deck.Id,
			Shuffled:  deck.Shuffled,
			Remaining: len(deck.Cards),
		})
	}
}
