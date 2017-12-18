
#Structure Below

Play some poker :-) `go run main.go`



# deck
`import "github.com/marcsantiago/go-poker/deck"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Card](#Card)
* [type Deck](#Deck)
  * [func (d *Deck) CreateDeck()](#Deck.CreateDeck)
  * [func (d *Deck) Draw() Card](#Deck.Draw)
  * [func (d *Deck) Shuffle()](#Deck.Shuffle)
* [type Hand](#Hand)
  * [func (slice Hand) Len() int](#Hand.Len)
  * [func (slice Hand) Less(i, j int) bool](#Hand.Less)
  * [func (slice Hand) Swap(i, j int)](#Hand.Swap)


#### <a name="pkg-files">Package files</a>
[colors.go](/src/github.com/marcsantiago/go-poker/deck/colors.go) [deck.go](/src/github.com/marcsantiago/go-poker/deck/deck.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    Spade   = color.New(color.FgBlue, color.Bold).SprintFunc()
    Clubs   = color.New(color.FgHiBlue, color.Bold).SprintFunc()
    Heart   = color.New(color.FgHiRed, color.Bold).SprintFunc()
    Diamond = color.New(color.FgRed, color.Bold).SprintFunc()
    Royal   = color.New(color.FgHiYellow, color.Bold).SprintFunc()
)
```



## <a name="Card">type</a> [Card](/src/target/deck.go?s=82:201#L9)
``` go
type Card struct {
    Suit      string
    Value     int
    IsRoyal   bool
    RoyalType string
    IsAce     bool
    Remove    bool
}
```
Card defines a cards basic attr










## <a name="Deck">type</a> [Deck](/src/target/deck.go?s=518:552#L26)
``` go
type Deck struct {
    Cards []Card
}
```
Deck is a collection is cards










### <a name="Deck.CreateDeck">func</a> (\*Deck) [CreateDeck](/src/target/deck.go?s=572:599#L31)
``` go
func (d *Deck) CreateDeck()
```
CreateDeck ...




### <a name="Deck.Draw">func</a> (\*Deck) [Draw](/src/target/deck.go?s=1343:1369#L73)
``` go
func (d *Deck) Draw() Card
```
Draw picks a card from the deck. If the deck is empty it creates a new deck and shuffles




### <a name="Deck.Shuffle">func</a> (\*Deck) [Shuffle](/src/target/deck.go?s=1527:1551#L85)
``` go
func (d *Deck) Shuffle()
```
Shuffle randomizes the deck order




## <a name="Hand">type</a> [Hand](/src/target/deck.go?s=243:259#L19)
``` go
type Hand []Card
```
Hand is a collection player of cards










### <a name="Hand.Len">func</a> (Hand) [Len](/src/target/deck.go?s=261:288#L21)
``` go
func (slice Hand) Len() int
```



### <a name="Hand.Less">func</a> (Hand) [Less](/src/target/deck.go?s=321:358#L22)
``` go
func (slice Hand) Less(i, j int) bool
```



### <a name="Hand.Swap">func</a> (Hand) [Swap](/src/target/deck.go?s=402:434#L23)
``` go
func (slice Hand) Swap(i, j int)
```










# player
`import "github.com/marcsantiago/go-poker/player"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CheckHand(hand deck.Hand) (string, int)](#CheckHand)
* [func ShowHand(p *Player)](#ShowHand)
* [type Player](#Player)
  * [func (p *Player) RemoveCards()](#Player.RemoveCards)


#### <a name="pkg-files">Package files</a>
[check_hand.go](/src/github.com/marcsantiago/go-poker/player/check_hand.go) [player.go](/src/github.com/marcsantiago/go-poker/player/player.go) [score_map.go](/src/github.com/marcsantiago/go-poker/player/score_map.go) 



## <a name="pkg-variables">Variables</a>
``` go
var Score = map[string]int{
    "royal flush":       200,
    "straight flush":    190,
    "full house":        170,
    "4 of a kind ace":   160,
    "4 of a kind king":  155,
    "4 of a kind queen": 150,
    "4 of a kind jack":  145,
    "4 of a kind 9":     139,
    "4 of a kind 8":     138,
    "4 of a kind 7":     137,
    "4 of a kind 6":     136,
    "4 of a kind 5":     135,
    "4 of a kind 4":     134,
    "4 of a kind 3":     133,
    "4 of a kind 2":     132,
    "4 of a kind 1":     131,

    "3 of a kind ace":   130,
    "3 of a kind king":  125,
    "3 of a kind queen": 120,
    "3 of a kind jack":  125,
    "3 of a kind 9":     119,
    "3 of a kind 8":     118,
    "3 of a kind 7":     117,
    "3 of a kind 6":     116,
    "3 of a kind 5":     115,
    "3 of a kind 4":     114,
    "3 of a kind 3":     113,
    "3 of a kind 2":     112,
    "3 of a kind 1":     111,

    "high pair of a kind ace":   110,
    "high pair of a kind king":  105,
    "high pair of a kind queen": 100,
    "high pair of a kind jack":  95,
    "high pair of a kind 9":     89,
    "high pair of a kind 8":     88,
    "high pair of a kind 7":     87,
    "high pair of a kind 6":     86,
    "high pair of a kind 5":     85,
    "high pair of a kind 4":     84,
    "high pair of a kind 3":     83,
    "high pair of a kind 2":     82,
    "high pair of a kind 1":     81,

    "pair of a kind ace":   80,
    "pair of a kind king":  79,
    "pair of a kind queen": 78,
    "pair of a kind jack":  77,
    "pair of a kind 9":     76,
    "pair of a kind 8":     75,
    "pair of a kind 7":     74,
    "pair of a kind 6":     73,
    "pair of a kind 5":     72,
    "pair of a kind 4":     71,
    "pair of a kind 3":     70,
    "pair of a kind 2":     69,
    "pair of a kind 1":     68,
}
```
Score maps a value to a hand type



## <a name="CheckHand">func</a> [CheckHand](/src/target/check_hand.go?s=136:180#L12)
``` go
func CheckHand(hand deck.Hand) (string, int)
```
CheckHand determines the hands value



## <a name="ShowHand">func</a> [ShowHand](/src/target/player.go?s=270:294#L19)
``` go
func ShowHand(p *Player)
```
ShowHand prints the hand to the screen




## <a name="Player">type</a> [Player](/src/target/player.go?s=131:226#L11)
``` go
type Player struct {
    Name       string
    TotalScore int
    HandsWon   int
    Cards      deck.Hand
}
```
Player holds informaation about the player










### <a name="Player.RemoveCards">func</a> (\*Player) [RemoveCards](/src/target/player.go?s=635:665#L35)
``` go
func (p *Player) RemoveCards()
```
RemoveCards creates a new hand with the cards removed








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
