package thegame

type GameRepository interface {
	GetCards() []Card
	GetDecks() []Deck
	NewDeck(deck *Deck) error
	GetDeck(id string) (Deck, error)
}
