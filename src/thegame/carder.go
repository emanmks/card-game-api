package thegame

import (
	"errors"
	"math/rand"
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

func (s *CardService) Shuffle(cards []Card, count uint) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	if count > uint(len(cards)) {
		return cards
	}

	return cards[0:count]
}

func (s *CardService) Draw(deckId string, count uint) (Draw, error) {
	deck, err := s.GetDeck(deckId)
	if err != nil {
		return Draw{}, errors.New("Can not draw the unknown deck")
	}

	var cards []Card

	if deck.Shuffled {
		cards = s.Shuffle(deck.Cards, count)
	} else {
		cards = deck.Cards[0:count]
	}

	newDraw := Draw{
		Deck:  deck,
		Cards: cards,
	}

	if s.repo.NewDraw(&newDraw) != nil {
		return Draw{}, errors.New("Failed creating a new deck")
	}

	deck.Cards = s.extractRemainingCards(deck.Cards, cards)
	if s.repo.UpdateDeck(deck) != nil {
		return Draw{}, errors.New("Failed updating the deck")
	}

	return newDraw, nil
}

func (s *CardService) extractRemainingCards(source []Card, target []Card) []Card {
	var remainingCards []Card

	for _, card := range source {
		if !s.contains(target, card) {
			remainingCards = append(remainingCards, card)
		}
	}

	return remainingCards
}

func (s *CardService) contains(cards []Card, card Card) bool {
	for _, v := range cards {
		if v == card {
			return true
		}
	}

	return false
}
