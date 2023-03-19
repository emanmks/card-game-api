package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"solaiman.me/cardgameapi/src/thegame"
)

type inMemoryRepository struct {
	cards []thegame.Card
	decks []thegame.Deck
}

func NewInMemoryRepository() thegame.GameRepository {
	repo := &inMemoryRepository{
		cards: []thegame.Card{
			{Value: "ACE SPADES", Suit: "SPADES", Code: "AS"},
			{Value: "2 SPADES", Suit: "SPADES", Code: "2S"},
			{Value: "3 SPADES", Suit: "SPADES", Code: "3S"},
			{Value: "4 SPADES", Suit: "SPADES", Code: "4S"},
			{Value: "5 SPADES", Suit: "SPADES", Code: "5S"},
			{Value: "6 SPADES", Suit: "SPADES", Code: "6S"},
			{Value: "7 SPADES", Suit: "SPADES", Code: "7S"},
			{Value: "8 SPADES", Suit: "SPADES", Code: "8S"},
			{Value: "9 SPADES", Suit: "SPADES", Code: "9S"},
			{Value: "10 SPADES", Suit: "SPADES", Code: "10S"},
			{Value: "JACK SPADES", Suit: "SPADES", Code: "JS"},
			{Value: "QUEEN SPADES", Suit: "SPADES", Code: "QS"},
			{Value: "KING SPADES", Suit: "SPADES", Code: "KS"},
			{Value: "ACE DIAMONDS", Suit: "DIAMONDS", Code: "AD"},
			{Value: "2 DIAMONDS", Suit: "DIAMONDS", Code: "2D"},
			{Value: "3 DIAMONDS", Suit: "DIAMONDS", Code: "3D"},
			{Value: "4 DIAMONDS", Suit: "DIAMONDS", Code: "4D"},
			{Value: "5 DIAMONDS", Suit: "DIAMONDS", Code: "5D"},
			{Value: "6 DIAMONDS", Suit: "DIAMONDS", Code: "6D"},
			{Value: "7 DIAMONDS", Suit: "DIAMONDS", Code: "7D"},
			{Value: "8 DIAMONDS", Suit: "DIAMONDS", Code: "8D"},
			{Value: "9 DIAMONDS", Suit: "DIAMONDS", Code: "9D"},
			{Value: "10 DIAMONDS", Suit: "DIAMONDS", Code: "10D"},
			{Value: "JACK DIAMONDS", Suit: "DIAMONDS", Code: "JD"},
			{Value: "QUEEN DIAMONDS", Suit: "DIAMONDS", Code: "QD"},
			{Value: "KING DIAMONDS", Suit: "DIAMONDS", Code: "KD"},
			{Value: "ACE CLUBS", Suit: "CLUBS", Code: "AC"},
			{Value: "2 CLUBS", Suit: "CLUBS", Code: "2C"},
			{Value: "3 CLUBS", Suit: "CLUBS", Code: "3C"},
			{Value: "4 CLUBS", Suit: "CLUBS", Code: "4C"},
			{Value: "5 CLUBS", Suit: "CLUBS", Code: "5C"},
			{Value: "6 CLUBS", Suit: "CLUBS", Code: "6C"},
			{Value: "7 CLUBS", Suit: "CLUBS", Code: "7C"},
			{Value: "8 CLUBS", Suit: "CLUBS", Code: "8C"},
			{Value: "9 CLUBS", Suit: "CLUBS", Code: "9C"},
			{Value: "10 CLUBS", Suit: "CLUBS", Code: "10C"},
			{Value: "JACK CLUBS", Suit: "CLUBS", Code: "JC"},
			{Value: "QUEEN CLUBS", Suit: "CLUBS", Code: "QC"},
			{Value: "KING CLUBS", Suit: "CLUBS", Code: "KC"},
			{Value: "ACE HEARTS", Suit: "HEARTS", Code: "AH"},
			{Value: "2 HEARTS", Suit: "HEARTS", Code: "2H"},
			{Value: "3 HEARTS", Suit: "HEARTS", Code: "3H"},
			{Value: "4 HEARTS", Suit: "HEARTS", Code: "4H"},
			{Value: "5 HEARTS", Suit: "HEARTS", Code: "5H"},
			{Value: "6 HEARTS", Suit: "HEARTS", Code: "6H"},
			{Value: "7 HEARTS", Suit: "HEARTS", Code: "7H"},
			{Value: "8 HEARTS", Suit: "HEARTS", Code: "8H"},
			{Value: "9 HEARTS", Suit: "HEARTS", Code: "9H"},
			{Value: "10 HEARTS", Suit: "HEARTS", Code: "10H"},
			{Value: "JACK HEARTS", Suit: "HEARTS", Code: "JH"},
			{Value: "QUEEN HEARTS", Suit: "HEARTS", Code: "QH"},
			{Value: "KING HEARTS", Suit: "HEARTS", Code: "KH"},
		},
		decks: []thegame.Deck{},
	}

	return repo
}

func (r *inMemoryRepository) GetCards() []thegame.Card {
	return r.cards
}

func (r *inMemoryRepository) GetDecks() []thegame.Deck {
	return r.decks
}

func (r *inMemoryRepository) NewDeck(deck *thegame.Deck) error {
	newDeck := thegame.Deck{
		Id:       uuid.New().String(),
		Shuffled: deck.Shuffled,
		Cards:    deck.Cards,
	}

	r.decks = append(r.decks, newDeck)
	deck.Id = newDeck.Id

	return nil
}

func (r *inMemoryRepository) GetDeck(id string) (thegame.Deck, error) {
	for _, deck := range r.decks {
		if deck.Id == id {
			return deck, nil
		}
	}

	return thegame.Deck{}, errors.New("Deck is not found")
}
