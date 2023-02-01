package poker

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Category int

const (
	StraightFlush Category = iota
	FourOfKind    Category = iota
	FullHouse     Category = iota
	Flush         Category = iota
	Straight      Category = iota
	ThreeOfKind   Category = iota
	TwoPair       Category = iota
	OnePair       Category = iota
	HighCard      Category = iota
)

type Card struct {
	rank int
	suit rune
}

func NewCard(input string) (Card, error) {
	card := []rune(input)
	if len(card) < 2 {
		return Card{}, fmt.Errorf("invalid card %s", input)
	}

	var suit rune

	switch card[len(card)-1] {
	case '♢', '♡', '♤', '♧':
		suit = card[len(card)-1]
	default:
		return Card{}, fmt.Errorf("invalid card %s", input)

	}

	switch rank := string(card[0 : len(card)-1]); {
	case rank == "1":
		return Card{}, fmt.Errorf("invalid rank %s", rank)
	case rank == "J":
		return Card{rank: 11, suit: suit}, nil
	case rank == "Q":
		return Card{rank: 12, suit: suit}, nil
	case rank == "K":
		return Card{rank: 13, suit: suit}, nil
	case rank == "A":
		return Card{rank: 1, suit: suit}, nil
	default:
		rank, err := strconv.Atoi(rank) // rank is now an int
		if err != nil {
			return Card{}, fmt.Errorf("invalid rank %d", rank)
		}
		if rank > 10 {
			return Card{}, fmt.Errorf("invalid rank %d", rank)
		}
		return Card{rank: rank, suit: suit}, nil
	}

}

type Hand []Card

func (h Hand) IsValid() bool {
	return len(h) != 5
}

func (h Hand) Rank() Category {
	sort.Slice(h, func(a, b int) bool { return h[a].rank < h[b].rank })

	suitGroup := map[rune][]int{}
	for _, card := range h {
		suitGroup[card.suit] = append(suitGroup[card.suit], card.rank)
	}

	unitGroup := map[int][]rune{}
	for _, card := range h {
		unitGroup[card.rank] = append(unitGroup[card.rank], card.suit)
	}

	if h[0].rank+4 == h[4].rank {
		if len(suitGroup) == 1 {
			return StraightFlush
		}
		return Straight
	}

	// Four of a kind : carre
	// Full house : brelan + pair
	// Straight: suite peu importe la couleur
	// Three of a kind : brelan
	// Two pair
	// One pair
	// High card

	return HighCard
}

func (h Hand) Value() (value int) {
	for _, card := range h {
		value += card.rank
	}
	return
}

func BestHand(input []string) ([]string, error) {

	hands := make([]Hand, 0)

	for _, h := range input {
		hand := make(Hand, 0)

		for _, c := range strings.Split(h, " ") {
			card, err := NewCard(c)
			if err != nil {
				return nil, err
			}

			hand = append(hand, card)
		}

		if hand.IsValid() {
			return nil, fmt.Errorf("wrong card number %d", len(hand))
		}

		hands = append(hands, hand)
	}

	return nil, nil
}
