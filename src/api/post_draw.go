package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"solaiman.me/cardgameapi/src/thegame"
)

type DrawRequestBody struct {
	DeckId string `json:"deck_id"`
	Count  uint   `json:"count"`
}

type NewDrawResponse struct {
	Id    string         `json:"id"`
	Cards []thegame.Card `json:"cards"`
}

func PostDrawHandler(handler RequestHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data DrawRequestBody
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newDraw, err := handler.cardService.Draw(data.DeckId, data.Count)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The draw was failed"})
			return
		}

		c.JSON(http.StatusOK, NewDrawResponse{
			Id:    newDraw.Id,
			Cards: newDraw.Cards,
		})
	}
}
