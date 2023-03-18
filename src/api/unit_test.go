package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"solaiman.me/cardgameapi/src/thegame"
)

func TestCanCreateANewRequestHandler(t *testing.T) {
	newHandler := NewHandler(thegame.CreateCardService())

	assert.NotEmpty(t, &newHandler.cardService, "The created handler has card service")
}
