package api

import "solaiman.me/cardgameapi/src/thegame"

type RequestHandler struct {
	cardService thegame.CardService
}

func NewHandler(cardService thegame.CardService) RequestHandler {
	return RequestHandler{
		cardService: cardService,
	}
}
