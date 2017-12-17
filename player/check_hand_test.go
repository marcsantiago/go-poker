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
			name: "Should be Royal Flush 1",
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
			name: "Should be Royal Flush ",
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
