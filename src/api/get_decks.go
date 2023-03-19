package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeckCollectionResponse struct {
	Id        string `json:"id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func GetDecksHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var deckCollection []DeckCollectionResponse
		decks := handler.cardService.GetDecks()
		for _, deck := range decks {
			deckCollection = append(deckCollection, DeckCollectionResponse{
				Id:        deck.Id,
				Shuffled:  deck.Shuffled,
				Remaining: len(deck.Cards),
			})
		}
		c.JSON(http.StatusOK, deckCollection)
	}
}
