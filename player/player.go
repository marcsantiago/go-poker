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
// THIS LOGIC MAY NOT WORK AS EXPECTED NEED TO WRITE TESTS
func (p *Player) RemoveCard(selected []int) {
	var cardsToRemove []deck.Card
	for _, s := range selected {
		cardsToRemove = append(cardsToRemove, p.Cards[s-1])
	}

	for _, card := range cardsToRemove {
		deleted := 0
		for i := range p.Cards {
			j := i - deleted
			if p.Cards[j] == card {
				p.Cards = p.Cards[:j+copy(p.Cards[j:], p.Cards[j+1:])]
				deleted++
			}
		}
	}
}
