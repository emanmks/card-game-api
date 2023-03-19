package thegame

import (
	"strings"
)

type CardService struct {
	repo GameRepository
}

func CreateCardService(repo GameRepository) CardService {
	return CardService{
		repo: repo,
	}
}

func (s *CardService) LoadAll() []Card {
	return s.repo.GetCards()
}

func (s *CardService) FilterCard(cards []Card, codes string) []Card {
	var filteredCards []Card

	for _, card := range cards {
		if strings.Contains(codes, card.Code) {
			filteredCards = append(filteredCards, card)
		}
	}

	return filteredCards
}

func (s *CardService) GetDecks() []Deck {
	return s.repo.GetDecks()
}

func (s *CardService) CreateNewDeck(shuffled bool, cardCodes string) Deck {
	var cards []Card

	if len(cardCodes) > 0 {
		cards = s.FilterCard(s.repo.GetCards(), cardCodes)
	} else {
		cards = s.repo.GetCards()
	}

	newDeck := Deck{
		Shuffled: shuffled,
		Cards:    cards,
	}

	s.repo.NewDeck(&newDeck)

	return newDeck
}

func (s *CardService) GetDeck(deckId string) (Deck, error) {
	return s.repo.GetDeck(deckId)
}
