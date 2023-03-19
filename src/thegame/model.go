package thegame

type Card struct {
	Value string `json:"value"`
	Code  string `json:"code"`
	Suit  string `json:"suit"`
}

type Deck struct {
	Id       string `json:"id"`
	Shuffled bool   `json:"shuffled"`
	Cards    []Card `json:"cards"`
}
