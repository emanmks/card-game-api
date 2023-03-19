package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"solaiman.me/cardgameapi/src/api"
	"solaiman.me/cardgameapi/src/repository/inmemory"
	"solaiman.me/cardgameapi/src/thegame"
)

func TestCanCreateANewRequestHandler(t *testing.T) {
	newHandler := api.NewHandler(thegame.CreateCardService(inmemory.NewInMemoryRepository()))

	assert.IsType(t, api.RequestHandler{}, newHandler, "The created handler has card service")
}
