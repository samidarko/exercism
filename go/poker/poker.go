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

type Cards []Card

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

	switch name := string(card[0 : len(card)-1]); {
	case name == "1":
		return Card{}, fmt.Errorf("invalid name %s", name)
	case name == "J":
		return Card{rank: 11, suit: suit}, nil
	case name == "Q":
		return Card{rank: 12, suit: suit}, nil
	case name == "K":
		return Card{rank: 13, suit: suit}, nil
	case name == "A":
		return Card{rank: 1, suit: suit}, nil
	default:
		rank, err := strconv.Atoi(name) // name is now an int
		if err != nil {
			return Card{}, fmt.Errorf("invalid name %d", rank)
		}
		if rank > 10 {
			return Card{}, fmt.Errorf("invalid name %d", rank)
		}
		return Card{rank: rank, suit: suit}, nil
	}

}

func sum(elements []int) (result int) {
	for _, element := range elements {
		result += element
	}
	return
}

func areLess(a, b [][]int) bool {
	for i := 1; i < len(a); i++ {
		if a[i][0] != b[i][0] {
			return a[1][0] < b[1][0]
		}
	}
	return false
}

type Hand struct {
	cards        []Card
	input        string
	category     Category
	combinations [][]int
}

func (h Hand) Less(other Hand) bool {

	if h.category == other.category {
		switch h.category {
		case Straight, StraightFlush, Flush:
			return sum(h.combinations[0]) < sum(other.combinations[0])
		case FourOfKind, FullHouse, TwoPair:
			hScore, otherScore := sum(h.combinations[0]), sum(other.combinations[0])
			if hScore == otherScore {
				return h.combinations[1][0] < other.combinations[1][0]
			}
			return hScore < otherScore
		case ThreeOfKind, OnePair:
			hScore, otherScore := sum(h.combinations[0]), sum(other.combinations[0])
			if hScore == otherScore {
				return areLess(h.combinations[1:], other.combinations[1:])
			}
			return hScore < otherScore

		case HighCard:
			return areLess(h.combinations, other.combinations)
		}
	}
	return h.category < other.category
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

	category, combinations := getCategory(hand.cards)
	hand.category = category
	hand.combinations = combinations

	return hand, nil
}

func getCategory(cards Cards) (Category, [][]int) {

	sort.Slice(cards, func(a, b int) bool { return cards[a].rank < cards[b].rank })

	firstCard := cards[0]
	isStraight := true
	isSameSuit := true
	combinations := [][]int{{firstCard.rank}}

	for i := 1; i < len(cards); i++ {
		if isStraight && cards[i-1].rank != cards[i].rank {
			isStraight = false
		}
		if isSameSuit && cards[i-1].suit != cards[i].suit {
			isSameSuit = false
		}
		if combinations[len(combinations)-1][0] == cards[i].rank {
			combinations[len(combinations)-1] = append(combinations[len(combinations)-1], cards[i].rank)
		} else {
			combinations = append(combinations, []int{cards[i].rank})
		}
	}

	sort.Slice(combinations, func(a, b int) bool {
		if len(combinations[a]) == len(combinations[b]) {
			return combinations[a][0] > combinations[b][0]
		}
		return len(combinations[a]) > len(combinations[b])
	})

	if isStraight {
		if isSameSuit {
			return StraightFlush, combinations
		}
		return Straight, combinations
	}

	if isSameSuit {
		return Flush, combinations
	}

	pairs := 0
	category := HighCard

	for _, streak := range combinations {
		switch len(streak) {
		case 4:
			category = FourOfKind
		case 3:
			category = ThreeOfKind
		case 2:
			pairs++
		}
	}

	switch category {
	case FourOfKind:
		return category, combinations
	case ThreeOfKind:
		if pairs == 1 {
			return FullHouse, combinations
		}
		return ThreeOfKind, combinations
	default:
		if pairs == 2 {
			return TwoPair, combinations
		} else if pairs == 1 {
			return OnePair, combinations
		} else {
			return HighCard, combinations
		}
	}
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
		return !hands[a].Less(hands[b])
	})

	highestCard := hands[0]

	result := make([]string, 0)

	for _, hand := range hands {
		if hand.Less(highestCard) {
			break
		}
		result = append(result, hand.input)
	}

	return result, nil
}
