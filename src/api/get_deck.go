package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"solaiman.me/cardgameapi/src/thegame"
)

type DeckParam struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type DeckResponse struct {
	Id        string         `json:"id"`
	Shuffled  bool           `json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     []thegame.Card `json:"cards"`
}

func GetDeckHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var deckParam DeckParam
		if err := c.ShouldBindUri(&deckParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deck id"})
			return
		}
		log.Println(fmt.Sprintf("Given deck ID: %s", deckParam.ID))
		deck, err := handler.cardService.GetDeck(deckParam.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The Deck is not found"})
			return
		}
		c.JSON(http.StatusOK, DeckResponse{
			Id:        deck.Id,
			Shuffled:  deck.Shuffled,
			Remaining: len(deck.Cards),
			Cards:     deck.Cards,
		})
	}
}
