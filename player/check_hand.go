package player

import (
	"sort"
	"strconv"

	"github.com/marcsantiago/go-poker/deck"
)

func CheckHand(hand deck.Hand) (string, int) {
	var tot int
	// sort hand for easy checking
	sort.Sort(hand)

	if isRoyalFlush(hand) {
		return "royal flush", 0
	}

	if isFlush(hand) {
		return "flush", 0
	}

	if isStraight(hand) {
		return "straight", 0
	}

	if isFullHouse(hand) {
		return "full house", 0
	}

	if val, ok := isFourOfAKind(hand); ok {
		return "4 of a kind " + val, 0
	}

	if val, ok := isThreeOfAKind(hand); ok {
		return "3 of a kind " + val, 0
	}

	if val, ok := isTwoPair(hand); ok {
		return "high pair of a kind " + val, 0
	}

	if val, ok := isPair(hand); ok {
		return "pair of a kind  " + val, 0
	}

	return "", tot
}

func isRoyalFlush(hand deck.Hand) bool {
	var nine, ten, jack, queen, king, ace bool

	if !isFlush(hand) {
		return false
	}

	for _, c := range hand {
		if c.IsRoyal {
			switch c.RoyalType {
			case deck.Royal("jack"):
				jack = true
			case deck.Royal("queen"):
				queen = true
			case deck.Royal("king"):
				king = true
			}
		}

		if c.IsAce {
			ace = true
			continue
		}

		if c.Value == 10 {
			ten = true
			continue
		}

		if c.Value == 9 {
			nine = true
			continue
		}
	}

	if nine && ten && jack && queen && king {
		return true
	}

	if ten && jack && queen && king && ace {
		return true
	}

	return false
}

// a flush is characterized as a hand of all the same suit
func isFlush(hand deck.Hand) bool {
	for i := range hand {
		j := i + 1
		if j < len(hand) {
			if hand[i].Suit != hand[j].Suit {
				return false
			}
		}
	}
	return true
}

func isStraight(hand deck.Hand) bool {
	for _, c := range hand {
		if c.IsRoyal {
			switch c.RoyalType {
			case deck.Royal("jack"):
				c.Value = 10
			case deck.Royal("queen"):
				c.Value = 11
			case deck.Royal("king"):
				c.Value = 12
			}
		}
	}

	sort.Sort(hand)
	for i := range hand {
		j := i + 1
		if j < len(hand) {
			if hand[i].Value+1 != hand[j].Value {
				return false
			}
		}
	}

	return true
}

func isFullHouse(hand deck.Hand) bool {
	nm := reduce(hand)
	if len(nm) > 2 || len(nm) < 2 {
		return false
	}

	var three, two bool
	for _, v := range nm {
		if v == 3 {
			three = true
			continue
		}
		if v == 2 {
			two = true
			continue
		}
	}

	return three && two
}

func isFourOfAKind(hand deck.Hand) (string, bool) {
	nm := reduce(hand)

	if len(nm) > 2 {
		return "", false
	}

	for k, v := range nm {
		if v == 4 {
			return getValue(k), true
		}
	}

	return "", false
}

func isThreeOfAKind(hand deck.Hand) (string, bool) {
	nm := reduce(hand)

	if len(nm) > 2 {
		return "", false
	}

	for k, v := range nm {
		if v == 3 {
			return getValue(k), true
		}
	}

	return "", false
}
func isTwoPair(hand deck.Hand) (string, bool) {
	nm := reduce(hand)
	var counter int

	var h int
	for k, v := range nm {
		if v == 2 {
			counter++
		}
		if k > h {
			h = k
		}
	}
	return getValue(h), (counter == 2)
}

func isPair(hand deck.Hand) (string, bool) {
	nm := reduce(hand)
	for k, v := range nm {
		if v == 2 {
			return getValue(k), true
		}
	}

	return "", false
}

func reduce(hand deck.Hand) map[int]int {
	for _, c := range hand {
		if c.IsRoyal {
			switch c.RoyalType {
			case deck.Royal("jack"):
				c.Value = 11
			case deck.Royal("queen"):
				c.Value = 12
			case deck.Royal("king"):
				c.Value = 13
			}
		}

		if c.IsAce {
			c.Value = 14
		}
	}

	nm := make(map[int]int)
	for _, c := range hand {
		nm[c.Value]++
	}
	return nm
}

func getValue(v int) string {
	switch v {
	case 11:
		return "jack"
	case 12:
		return "queen"
	case 13:
		return "king"
	case 14:
		return "ace"
	default:
		return strconv.Itoa(v)
	}
}
