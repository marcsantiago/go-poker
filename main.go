package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/marcsantiago/go-poker/deck"
	"github.com/marcsantiago/go-poker/player"
)

var (
	playerOne player.Player
	computer  player.Player
)

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

func computerLogic(comp *player.Player, deck *deck.Deck) (string, int) {
	// first check the hand to see of the score is high enough to do nothing
	// computer will stay with a hand of 3 of a kind or greater
	handName, baseScore := player.CheckHand(comp.Cards)
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
		source := rand.NewSource(t.UnixNano())
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

	for _, s := range selected {
		comp.Cards[s-1].Remove = true
	}

	comp.RemoveCards()

	for i := 0; i < len(selected); i++ {
		comp.Cards = append(comp.Cards, deck.Draw())
	}
	return player.CheckHand(comp.Cards)
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
					for _, s := range selected {
						playerOne.Cards[s-1].Remove = true
					}
					playerOne.RemoveCards()
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

		player1HandName, baseScore := player.CheckHand(playerOne.Cards)
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
