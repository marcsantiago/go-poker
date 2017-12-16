package deck

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Spade   = color.New(color.FgBlue, color.Bold).SprintFunc()
	Clubs   = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Heart   = color.New(color.FgHiRed, color.Bold).SprintFunc()
	Diamond = color.New(color.FgRed, color.Bold).SprintFunc()
	Royal   = color.New(color.FgHiYellow, color.Bold).SprintFunc()

	suit      = []string{fmt.Sprintf("%s", Heart("hearts")), fmt.Sprintf("%s", Spade("spades")), fmt.Sprintf("%s", Clubs("clubs")), fmt.Sprintf("%s", Diamond("diamonds"))}
	royalSuit = []string{fmt.Sprintf("%s", Royal("queen")), fmt.Sprintf("%s", Royal("jack")), fmt.Sprintf("%s", Royal("king"))}
)
