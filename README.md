
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







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
