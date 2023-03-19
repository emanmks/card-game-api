package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"solaiman.me/cardgameapi/src/thegame"
)

type HealthCheck struct {
	Ack string `json:"ack"`
}

type NewDeckResponse struct {
	Id        string `json:"id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type DeckDetailsResponse struct {
	Id        string         `json:"id"`
	Shuffled  bool           `json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     []thegame.Card `json:"cards"`
}

type DrawResponse struct {
	Cards []thegame.Card `json:"cards"`
}

func TestHealthCheckEndpoint(t *testing.T) {
	router := setupRouter()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(rr, req)

	var healthCheck HealthCheck
	json.Unmarshal(rr.Body.Bytes(), &healthCheck)

	assert.Equal(t, http.StatusOK, rr.Code, "Health check status should be OK")
	assert.NotEmpty(t, healthCheck.Ack, "Ack attribute should be available in the health check response")
}

func TestCardsRoute(t *testing.T) {
	router := setupRouter()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/cards", nil)
	router.ServeHTTP(rr, req)

	var cards []thegame.Card
	json.Unmarshal(rr.Body.Bytes(), &cards)

	assert.Equal(t, http.StatusOK, rr.Code, "Http status code should be OK")
	for _, card := range cards {
		assert.NotEmpty(t, card.Code, "Card Code is available")
		assert.NotEmpty(t, card.Suit, "Card Suit is available")
		assert.NotEmpty(t, card.Value, "Card Value is available")
	}
}

func TestPostDeckEndpointTest(t *testing.T) {
	router := setupRouter()

	rr := httptest.NewRecorder()
	jsonParam := `{"shuffled":false}`
	req, _ := http.NewRequest(http.MethodPost, "/deck", strings.NewReader(string(jsonParam)))
	router.ServeHTTP(rr, req)

	var newDeckResponse NewDeckResponse
	json.Unmarshal(rr.Body.Bytes(), &newDeckResponse)

	assert.Equal(t, http.StatusOK, rr.Code, "Http status code should be OK")
	assert.NotEmpty(t, newDeckResponse.Id, "Deck Id should be available in the response")
	assert.IsType(t, false, newDeckResponse.Shuffled, "Deck Shuffled flag should be available in the response")
	assert.NotEmpty(t, newDeckResponse.Remaining, "Deck Remaining should be available in the response")
}

func TestGetDeckEndpointTest(t *testing.T) {
	router := setupRouter()

	postrr := httptest.NewRecorder()

	jsonParam := `{"shuffled":false}`
	postReq, _ := http.NewRequest(http.MethodPost, "/deck", strings.NewReader(string(jsonParam)))
	router.ServeHTTP(postrr, postReq)

	var newDeckResponse NewDeckResponse
	json.Unmarshal(postrr.Body.Bytes(), &newDeckResponse)

	getrr := httptest.NewRecorder()
	findUri := "/deck/" + newDeckResponse.Id
	req, _ := http.NewRequest(http.MethodGet, findUri, nil)
	router.ServeHTTP(getrr, req)

	var deckResponse DeckDetailsResponse
	json.Unmarshal(getrr.Body.Bytes(), &deckResponse)

	assert.Equal(t, http.StatusOK, getrr.Code, "Http status code should be OK")
	assert.NotEmpty(t, deckResponse.Id, "Deck Id should be available in the response")
	assert.IsType(t, false, deckResponse.Shuffled, "Deck Shuffled flag should be available in the response")
	assert.NotEmpty(t, deckResponse.Remaining, "Deck Remaining should be available in the response")
	assert.NotEmpty(t, deckResponse.Cards, "Deck Remaining Cards should be available in the response")
}

func TestPostDrawEndpointTest(t *testing.T) {
	router := setupRouter()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/draw", nil)
	router.ServeHTTP(rr, req)

	var drawResponse DrawResponse
	json.Unmarshal(rr.Body.Bytes(), &drawResponse)

	assert.Equal(t, http.StatusOK, rr.Code, "Http status code should be OK")
	assert.NotEmpty(t, drawResponse.Cards, "Cards should be available in the response")
}
