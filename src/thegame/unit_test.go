package thegame

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanProvideListOfCards(t *testing.T) {
	service := CreateCardService()

	cards := service.LoadAll()

	for _, card := range cards {
		assert.NotEmpty(t, card.Code, "Card Code is available")
		assert.NotEmpty(t, card.Suit, "Card Suit is available")
		assert.NotEmpty(t, card.Value, "Card Value is available")
	}
}

func TestCanFilterCardsByListOfCodesInAString(t *testing.T) {
	service := CreateCardService()

	codes := "AS,2S"
	filteredCards := service.FilterCard(service.cards, codes)

	for _, card := range filteredCards {
		assert.True(t, strings.Contains(codes, card.Code), "Card Value is found in the code list string")
	}
}
