package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"solaiman.me/cardgameapi/src/api"
	"solaiman.me/cardgameapi/src/repository/inmemory"
	"solaiman.me/cardgameapi/src/thegame"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	requestHandler := api.NewHandler(thegame.CreateCardService(inmemory.NewInMemoryRepository()))

	// Ping test
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ack": time.Now(),
		})
	})

	// Get card list
	r.GET("/cards", api.GetCardsHandler(requestHandler))

	// Get Deck list
	r.GET("/decks", api.GetDecksHandler(requestHandler))

	// POST a new deck
	r.POST("/deck", api.PostDeckHandler(requestHandler))

	// Get a single deck
	r.GET("/deck/:id", api.GetDeckHandler(requestHandler))

	r.POST("/draw", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cards": []thegame.Card{
				{
					Value: "ACE",
					Suit:  "SPADES",
					Code:  "AS",
				},
				{
					Value: "KING",
					Suit:  "HEARTS",
					Code:  "KH",
				},
				{
					Value: "8",
					Suit:  "CLUBS",
					Code:  "8C",
				},
			},
		})
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
