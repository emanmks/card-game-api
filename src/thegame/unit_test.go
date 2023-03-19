package thegame_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"solaiman.me/cardgameapi/src/repository/inmemory"
	"solaiman.me/cardgameapi/src/thegame"
)

func TestCanProvideListOfCards(t *testing.T) {
	service := thegame.CreateCardService(inmemory.NewInMemoryRepository())

	cards := service.LoadAll()

	for _, card := range cards {
		assert.NotEmpty(t, card.Code, "Card Code is available")
		assert.NotEmpty(t, card.Suit, "Card Suit is available")
		assert.NotEmpty(t, card.Value, "Card Value is available")
	}
}

func TestCanFilterCardsByListOfCodesInAString(t *testing.T) {
	service := thegame.CreateCardService(inmemory.NewInMemoryRepository())

	codes := "AS,2S"
	filteredCards := service.FilterCard(service.LoadAll(), codes)

	for _, card := range filteredCards {
		assert.True(t, strings.Contains(codes, card.Code), "Card Value is found in the code list string")
	}
}
