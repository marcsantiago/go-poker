package player

import (
	"fmt"
	"sort"

	"github.com/marcsantiago/go-poker/deck"
)

// Player holds informaation about the player
type Player struct {
	Name       string
	TotalScore int
	HandsWon   int
	Cards      deck.Hand
}

// ShowHand prints the hand to the screen
func ShowHand(p *Player) {
	sort.Sort(p.Cards)
	for n, c := range p.Cards {
		if c.IsRoyal {
			fmt.Printf("%d) %v of %v\n", n+1, c.RoyalType, c.Suit)
		} else if c.IsAce {
			fmt.Printf("%d) %v of %v\n", n+1, "ace", c.Suit)
		} else {
			fmt.Printf("%d) %v of %v\n", n+1, c.Value, c.Suit)
		}

	}
}

// RemoveCard removes the card from the players hand
func (p *Player) RemoveCards() {
	newHand := make([]deck.Card, 0, 5)
	for _, c := range p.Cards {
		if !c.Remove {
			newHand = append(newHand, c)
		}
	}
	p.Cards = newHand
}
