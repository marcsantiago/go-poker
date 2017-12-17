package player

import (
	"testing"

	"github.com/marcsantiago/go-poker/deck"
)

func Test_isRoyalFlush(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be Royal Straight Flush 1",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{Value: 10},
				deck.Card{Value: 9},
			}},
			want: true,
		},
		{
			name: "Should be Royal Straight Flush 2",
			args: args{deck.Hand{
				deck.Card{IsAce: true},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{Value: 10},
			}},
			want: true,
		},
		{
			name: "Should not be a Royal Flush",
			args: args{deck.Hand{
				deck.Card{IsAce: true},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{Value: 10},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{Value: 10},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRoyalFlush(tt.args.hand); got != tt.want {
				t.Errorf("isRoyalFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFlush(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be a flush",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king"), Suit: deck.Heart("hearts")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen"), Suit: deck.Heart("hearts")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack"), Suit: deck.Heart("hearts")},
				deck.Card{Value: 10, Suit: deck.Heart("hearts")},
				deck.Card{Value: 9, Suit: deck.Heart("hearts")},
			}},
			want: true,
		},
		{
			name: "Should be a flush 2",
			args: args{deck.Hand{
				deck.Card{Value: 8, Suit: deck.Clubs("club")},
				deck.Card{Value: 9, Suit: deck.Clubs("club")},
				deck.Card{Value: 10, Suit: deck.Clubs("club")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack"), Suit: deck.Clubs("club")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen"), Suit: deck.Clubs("club")},
			}},
			want: true,
		},
		{
			name: "Should not be a flush",
			args: args{deck.Hand{
				deck.Card{Value: 6, Suit: deck.Clubs("club")},
				deck.Card{Value: 7, Suit: deck.Heart("hearts")},
				deck.Card{Value: 8, Suit: deck.Clubs("club")},
				deck.Card{Value: 9, Suit: deck.Clubs("club")},
				deck.Card{Value: 10, Suit: deck.Clubs("club")},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlush(tt.args.hand); got != tt.want {
				t.Errorf("isFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraight(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be a straight flush",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{Value: 10},
				deck.Card{Value: 9},
			}},
			want: true,
		},
		{
			name: "Should be a straight flush 2",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 10},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want: true,
		},
		{
			name: "Should be a straight flush 3",
			args: args{deck.Hand{
				deck.Card{Value: 6},
				deck.Card{Value: 7},
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 10},
			}},
			want: true,
		},
		{
			name: "Should not be a straight flush",
			args: args{deck.Hand{
				deck.Card{Value: 5},
				deck.Card{Value: 7},
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 10},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.hand); got != tt.want {
				t.Errorf("isFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFullHouse(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should be a full house",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{Value: 2},
				deck.Card{Value: 2},
			}},
			want: true,
		},
		{
			name: "Should be a full house 2",
			args: args{deck.Hand{
				deck.Card{Value: 3},
				deck.Card{Value: 3},
				deck.Card{Value: 2},
				deck.Card{Value: 2},
				deck.Card{Value: 2},
			}},
			want: true,
		},
		{
			name: "Should not be a full house",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 10},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFullHouse(tt.args.hand); got != tt.want {
				t.Errorf("isFullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFourOfAKind(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Should be a 4 of a kind",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "8",
			want1: true,
		},
		{
			name: "Should be a 4 of a kind 2",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "king",
			want1: true,
		},
		{
			name: "Should not be a 4 of a kind",
			args: args{deck.Hand{
				deck.Card{Value: 3},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isFourOfAKind(tt.args.hand)
			if got != tt.want {
				t.Errorf("isFourOfAKind() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isFourOfAKind() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isThreeOfAKind(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Should be a 3 of a kind",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 7},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "8",
			want1: true,
		},
		{
			name: "Should be a 3 of a kind 2",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "queen",
			want1: true,
		},
		{
			name: "Should not be a 3 of a kind",
			args: args{deck.Hand{
				deck.Card{Value: 3},
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 3},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isThreeOfAKind(tt.args.hand)
			if got != tt.want {
				t.Errorf("isThreeOfAKind() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isThreeOfAKind() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isTwoPair(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Should be a 2 pairs",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 7},
				deck.Card{Value: 7},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "8",
			want1: true,
		},
		{
			name: "Should be a 2 pairs 2",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "king",
			want1: true,
		},
		{
			name: "Should not be a 2 pairs",
			args: args{deck.Hand{
				deck.Card{Value: 3},
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 3},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "3",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isTwoPair(tt.args.hand)
			if got != tt.want {
				t.Errorf("isTwoPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isTwoPair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isPair(t *testing.T) {
	type args struct {
		hand deck.Hand
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Should be a pair",
			args: args{deck.Hand{
				deck.Card{Value: 8},
				deck.Card{Value: 8},
				deck.Card{Value: 6},
				deck.Card{Value: 4},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "8",
			want1: true,
		},
		{
			name: "Should be a pair 2",
			args: args{deck.Hand{
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("king")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("jack")},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
				deck.Card{Value: 4},
			}},
			want:  "king",
			want1: true,
		},
		{
			name: "Should not be a 2 pairs",
			args: args{deck.Hand{
				deck.Card{Value: 1},
				deck.Card{Value: 8},
				deck.Card{Value: 9},
				deck.Card{Value: 3},
				deck.Card{IsRoyal: true, RoyalType: deck.Royal("queen")},
			}},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isPair(tt.args.hand)
			if got != tt.want {
				t.Errorf("isPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isPair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
