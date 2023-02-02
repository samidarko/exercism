package poker

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Category int

const (
	HighCard      Category = iota
	OnePair       Category = iota
	TwoPair       Category = iota
	ThreeOfKind   Category = iota
	Straight      Category = iota
	Flush         Category = iota
	FullHouse     Category = iota
	FourOfKind    Category = iota
	StraightFlush Category = iota
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

type Hand struct {
	cards    []Card
	input    string
	category Category
}

func NewHand(input string) (Hand, error) {
	hand := Hand{
		input: input,
		cards: make([]Card, 0),
	}

	for _, c := range strings.Split(input, " ") {
		card, err := NewCard(c)
		if err != nil {
			return Hand{}, err
		}

		hand.cards = append(hand.cards, card)
	}

	if len(hand.cards) != 5 {
		return Hand{}, fmt.Errorf("wrong card number %d", len(hand.cards))
	}

	hand.category = getCategory(hand.cards)

	return hand, nil
}

func getCategory(cards []Card) Category {
	sort.Slice(cards, func(a, b int) bool { return cards[a].rank < cards[b].rank })

	suitGroup := map[rune][]int{}
	for _, card := range cards {
		suitGroup[card.suit] = append(suitGroup[card.suit], card.rank)
	}

	if len(suitGroup) == 1 {
		return Flush
	}

	if cards[0].rank+4 == cards[4].rank {
		if len(suitGroup) == 1 {
			return StraightFlush
		}
		return Straight
	}

	unitGroup := map[int][]rune{}
	for _, card := range cards {
		unitGroup[card.rank] = append(unitGroup[card.rank], card.suit)
	}

	pairs := 0
	threeOfKind := false

	for _, suits := range unitGroup {
		if len(suits) == 4 {
			return FourOfKind
		}
		if len(suits) == 3 {
			threeOfKind = true
		}
		if len(suits) == 2 {
			pairs++
		}
	}

	if threeOfKind {
		if pairs == 1 {
			return FullHouse
		}
		return ThreeOfKind
	}

	switch pairs {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) Value() (value int) {
	for _, card := range h.cards {
		switch h.category {
		case HighCard:
			if card.rank > value {
				value = card.rank
			}
		default:
			value += card.rank
		}
	}
	return
}

func BestHand(input []string) ([]string, error) {

	hands := make([]Hand, 0)

	for _, h := range input {

		hand, err := NewHand(h)
		if err != nil {
			return nil, err
		}

		hands = append(hands, hand)
	}

	sort.Slice(hands, func(a, b int) bool {
		aCategory, bCategory := hands[a].category, hands[b].category
		if aCategory == bCategory {
			return hands[a].Value() > hands[b].Value()
		}
		return aCategory > bCategory
	})

	highestCard := hands[0]

	result := make([]string, 0)

	for _, hand := range hands {
		if hand.category == highestCard.category && hand.Value() == highestCard.Value() {
			result = append(result, hand.input)
		}
	}

	return result, nil
}
