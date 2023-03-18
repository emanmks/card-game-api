package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
	Id        string `json:"id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type DrawResponse struct {
	Cards []Card `json:"cards"`
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

	var cards []Card
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
	req, _ := http.NewRequest(http.MethodPost, "/deck", nil)
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

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/deck/id", nil)
	router.ServeHTTP(rr, req)

	var deckResponse DeckDetailsResponse
	json.Unmarshal(rr.Body.Bytes(), &deckResponse)

	assert.Equal(t, http.StatusOK, rr.Code, "Http status code should be OK")
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
