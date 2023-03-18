package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"solaiman.me/cardgameapi/src/api"
	"solaiman.me/cardgameapi/src/thegame"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	requestHandler := api.NewHandler(thegame.CreateCardService())

	// Ping test
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ack": time.Now(),
		})
	})

	r.GET("/cards", api.GetCardsHandler(requestHandler))

	r.POST("/deck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":        "a251071b-662f-44b6-ba11-e24863039c59",
			"shuffled":  false,
			"remaining": 30,
		})
	})

	r.GET("/deck/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":        "a251071b-662f-44b6-ba11-e24863039c59",
			"shuffled":  false,
			"remaining": 3,
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
