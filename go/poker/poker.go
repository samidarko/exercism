package poker

import (
	"fmt"
	"strconv"
	"strings"
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

func BestHand(hands []string) ([]string, error) {

	for _, h := range hands {
		hand := make([]Card, 0)

		for _, c := range strings.Split(h, " ") {
			card, err := NewCard(c)
			if err != nil {
				return nil, err
			}

			hand = append(hand, card)
		}

		if len(hand) != 5 {
			return nil, fmt.Errorf("wrong card number %d", len(hand))
		}
	}

	return nil, nil
}
