package deck

import (
	"math/rand"
	"time"
)

// Card defines a cards basic attr
type Card struct {
	Suit      string
	Value     int
	IsRoyal   bool
	RoyalType string
	IsAce     bool
	Remove    bool
}

// Hand is a collection player of cards
type Hand []Card

func (slice Hand) Len() int           { return len(slice) }
func (slice Hand) Less(i, j int) bool { return slice[i].Value < slice[j].Value }
func (slice Hand) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

// Deck is a collection is cards
type Deck struct {
	Cards []Card
}

// CreateDeck ...
func (d *Deck) CreateDeck() {
	d.Cards = make([]Card, 0, 52)

	// add reg cards
	for i := 1; i < 10; i++ {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:      s,
				Value:     i,
				IsRoyal:   false,
				RoyalType: "",
				IsAce:     false,
			})
		}
	}

	for _, r := range royalSuit {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:      s,
				Value:     10,
				IsRoyal:   true,
				RoyalType: r,
				IsAce:     false,
			})
		}
	}

	for _, s := range suit {
		d.Cards = append(d.Cards, Card{
			Suit:      s,
			Value:     1,
			IsRoyal:   false,
			RoyalType: "",
			IsAce:     true,
		})
	}
}

// Draw picks a card from the deck. If the deck is empty it creates a new deck and shuffles
func (d *Deck) Draw() Card {
	if len(d.Cards) != 0 {
		card := d.Cards[0]
		d.Cards = d.Cards[1:]
		return card
	}
	d.CreateDeck()
	d.Shuffle()
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// Shuffle randomizes the deck order
func (d *Deck) Shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	arr := r.Perm(52)
	for i, a := range arr {
		d.Cards[i] = d.Cards[a]
	}
}
