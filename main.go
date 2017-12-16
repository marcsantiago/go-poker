package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/marcsantiago/go-poker/deck"
	"github.com/marcsantiago/go-poker/player"
)

var (
	playerOne player.Player
	computer  player.Player
)

func checkHand(hand deck.Hand) (string, int) {
	// check for royal flush
	sort.Sort(hand)
	numberOfTens := 0
	numberOfRoyals := 0
	numberOfAces := 0
	allSameSuit := true
	future := 1

	for i, c := range hand {
		if c.IsRoyal {
			numberOfRoyals++
		}
		if c.IsAce {
			numberOfAces++
		}
		if c.Value == 10 && !c.IsRoyal {
			numberOfTens++
		}
		if future < len(hand) {
			if hand[i].Suit != hand[future].Suit {
				allSameSuit = false
			}
		}
		future++
	}

	if allSameSuit {
		if numberOfRoyals == 3 { // jack, queen, king
			if numberOfAces == 1 { // ace
				if numberOfTens == 1 { // 10
					return "royal flush", 0
				}
			}
		}
	}

	// check for flush
	if allSameSuit {
		flush := true
		future = 1
		for i := range hand {
			if future < len(hand) {
				if hand[i].Value+1 != hand[future].Value {
					flush = false
					break
				}
			}
		}
		if flush {
			return "flush", 0
		}
	}

	// check for straight
	straight := true
	future = 1
	for i := range hand {
		if future < len(hand) {
			cardValue := hand[i].Value
			futureValue := hand[future].Value

			if hand[i].IsRoyal {
				if hand[i].RoyalType == deck.Royal("jack") {
					cardValue = 11
				} else if hand[i].RoyalType == deck.Royal("queen") {
					cardValue = 12
				} else if hand[i].RoyalType == deck.Royal("king") {
					cardValue = 13
				}
			}
			if hand[future].IsRoyal {
				if hand[i].RoyalType == deck.Royal("jack") {
					futureValue = 11
				} else if hand[i].RoyalType == deck.Royal("queen") {
					futureValue = 12
				} else if hand[i].RoyalType == deck.Royal("king") {
					futureValue = 13
				}
			}

			if cardValue+1 != futureValue {
				straight = false
				break
			}
		}
		future++
	}

	if straight {
		return "straight", 0
	}

	// full house?
	pair := false
	threeOfAKind := false

	for _, c := range hand {
		cards := findCardsByValue(hand, c.Value, c.IsRoyal, c.RoyalType)
		if len(cards) == 3 {
			threeOfAKind = true
		}
		if len(cards) == 2 {
			pair = true
		}
	}
	if pair && threeOfAKind {
		return "full house", 0
	}

	// determine if 4 of a kind, 3 or a kind, or 2 of a kind
	twoPairs := []string{}
	for _, card := range hand {
		cards := findCardsByValue(hand, card.Value, card.IsRoyal, card.RoyalType)

		if len(cards) == 4 {
			if cards[0].IsAce {
				return "4 of a kind ace", 0
			} else if cards[0].IsRoyal {
				return fmt.Sprintf("4 of a kind %s", cards[0].RoyalType), 0
			} else {
				return fmt.Sprintf("4 of a kind %d", cards[0].Value), 0
			}
		}
		if len(cards) == 3 {
			if cards[0].IsAce {
				return "3 of a kind ace", 0
			} else if cards[0].IsRoyal {
				return fmt.Sprintf("3 of a kind %s", cards[0].RoyalType), 0
			} else {
				return fmt.Sprintf("3 of a kind %d", cards[0].Value), 0
			}
		}
		if len(cards) == 2 {
			if cards[0].IsAce {
				if !inSlice(twoPairs, "pair of a kind ace") {
					twoPairs = append(twoPairs, "pair of a kind ace")
				}

			} else if cards[0].IsRoyal {
				if !inSlice(twoPairs, fmt.Sprintf("pair of a kind %s", cards[0].RoyalType)) {
					twoPairs = append(twoPairs, fmt.Sprintf("pair of a kind %s", cards[0].RoyalType))
				}
			} else {
				if !inSlice(twoPairs, fmt.Sprintf("pair of a kind %d", cards[0].Value)) {
					twoPairs = append(twoPairs, fmt.Sprintf("pair of a kind %d", cards[0].Value))
				}
			}
		}
	}
	if len(twoPairs) == 2 {
		if player.Score[twoPairs[0]] > player.Score[twoPairs[1]] {
			return "high " + twoPairs[0], 0
		}
		return "high " + twoPairs[1], 0
	} else if len(twoPairs) == 1 {
		return twoPairs[0], 0
	}

	// logic here is broken
	HighTwoCards := hand[len(hand)-2:]
	tot := 0
	for _, c := range HighTwoCards {
		if c.IsRoyal {
			if c.RoyalType == fmt.Sprintf("%s", deck.Royal("king")) {
				tot = tot + 13
			} else if c.RoyalType == fmt.Sprintf("%s", deck.Royal("queen")) {
				tot = tot + 12
			} else if c.RoyalType == fmt.Sprintf("%s", deck.Royal("jack")) {
				tot = tot + 11
			}
		} else {
			tot = tot + c.Value
		}
	}
	return "", tot
}

func computerLogic(comp *player.Player, deck *deck.Deck) (string, int) {
	// first check the hand to see of the score is high enough to do nothing
	// computer will stay with a hand of 3 of a kind or greater
	handName, baseScore := checkHand(comp.Cards)
	compScore := baseScore
	if baseScore > 0 {
		compScore = player.Score[handName]
	}
	if compScore >= 111 {
		return handName, 0
	}
	// determine how many and which cards to RemoveCard
	// later I will go back and add probalities of how many cards to discard to maximize the handName
	// for now just to get the game working I'm going to randomize 0-5
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	n := r.Perm(5)[0] // number of cards to drop

	selected := []int{}
	for i := 0; i < n; i++ {
		t := time.Now()
		source := rand.NewSource(t.Unix())
		r := rand.New(source)
		n := r.Perm(5)[0] + 1
		contains := func(n int, sel []int) bool {
			for _, s := range selected {
				if n == s {
					return true
				}
			}
			return false
		}(n, selected)
		if !contains {
			selected = append(selected, n)
		}
	}
	comp.RemoveCard(selected)
	for i := 0; i < len(selected); i++ {
		comp.Cards = append(comp.Cards, deck.Draw())
	}
	return checkHand(comp.Cards)
}

func findCardsByValue(hand deck.Hand, value int, isRoyal bool, royalType string) []deck.Card {
	var cards deck.Hand
	for _, c := range hand {
		if isRoyal {
			if c.RoyalType == royalType {
				cards = append(cards, c)
			}
		} else {
			if c.Value == value {
				cards = append(cards, c)
			}
		}

	}
	return cards
}

func inSlice(slice []string, tar string) bool {
	for _, s := range slice {
		if s == tar {
			return true
		}
	}
	return false
}

func initHand(cards *deck.Deck, p *player.Player, c *player.Player) {
	p.Cards = make([]deck.Card, 0, 5)
	c.Cards = make([]deck.Card, 0, 5)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			p.Cards = append(p.Cards, cards.Draw())
		} else {
			c.Cards = append(c.Cards, cards.Draw())
		}
	}
}

func main() {
	var d deck.Deck
	d.CreateDeck()
	d.Shuffle()
	// game loop
	for {
		initHand(&d, &playerOne, &computer) // alternating 5 card Draw
		fmt.Println("Player Draw 5 cards")
		player.ShowHand(&playerOne)

		var discard int
		for {
			fmt.Println("How many cards would like you to discard? (0-5)")
			_, err := fmt.Scan(&discard)
			if err != nil {
				fmt.Println("Please only enter a digit between 0 and 5 inclusive")
				continue
			}
			break
		}

		if discard > 0 {
			player.ShowHand(&playerOne)

			fmt.Println("Please choose which cards to discard (0-5)")
			var selected []int
			var input int
			for {

				if len(selected) == discard {
					playerOne.RemoveCard(selected)
					break
				}

				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Please only enter a digit between 1 and 5 inclusive")
					player.ShowHand(&playerOne)
					continue
				}
				selected = append(selected, input)
			}

			fmt.Printf("Picking up %d more cards\n", discard)
			for i := 0; i < discard; i++ {
				playerOne.Cards = append(playerOne.Cards, d.Draw())
			}
			fmt.Println("New hand")
			player.ShowHand(&playerOne)
			fmt.Printf("\n\n")
		}

		player1HandName, baseScore := checkHand(playerOne.Cards)
		player1Score := baseScore
		if player1Score == 0 {
			player1Score = player.Score[player1HandName]
		}
		playerOne.TotalScore = playerOne.TotalScore + player1Score

		computerHandName, baseScore := computerLogic(&computer, &d)
		computerScore := baseScore
		if computerScore == 0 {
			computerScore = player.Score[computerHandName]
		}
		computer.TotalScore = computer.TotalScore + computerScore

		if computerScore > player1Score {
			fmt.Printf("Computer won with the hand: %s score: %d\n", computerHandName, computerScore)
			computer.HandsWon++
		} else {
			fmt.Printf("Player1 won with the hand: %s score: %d\n", player1HandName, player1Score)
			playerOne.HandsWon++
		}

		// another game?
		for {
			fmt.Printf("play again? [y, n]\n")
			var action string
			fmt.Scan(&action)
			if strings.Contains(strings.ToLower(action), "y") {
				fmt.Printf("\n\n")
				break
			}
			fmt.Printf("Player1 final score: %d\nPlayer1 total hands won %d\n", playerOne.TotalScore, playerOne.HandsWon)
			fmt.Printf("Computer final score: %d\nComputer total hands won %d\n", computer.TotalScore, computer.HandsWon)
			return
		}
	}

}