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
		return Card{rank: 14, suit: suit}, nil
	default:
		rank, err := strconv.Atoi(name)
		if err != nil {
			return Card{}, fmt.Errorf("invalid name %d", rank)
		}
		if rank > 10 {
			return Card{}, fmt.Errorf("invalid name %d", rank)
		}
		return Card{rank: rank, suit: suit}, nil
	}

}

type Hand struct {
	input    string
	category Category
	ranks    []int
}

func (h Hand) Equal(other Hand) bool {
	if h.category != other.category {
		return false
	}
	for i := range h.ranks {
		if h.ranks[i] != other.ranks[i] {
			return false
		}
	}
	return true
}

func (h Hand) Less(other Hand) bool {

	if h.category == other.category {
		for i := range h.ranks {
			if h.ranks[i] != other.ranks[i] {
				return h.ranks[i] < other.ranks[i]
			}
		}
		return false

	}
	return h.category < other.category
}

func NewHand(input string) (Hand, error) {

	cards := make([]Card, 0)

	for _, c := range strings.Split(input, " ") {
		card, err := NewCard(c)
		if err != nil {
			return Hand{}, err
		}

		cards = append(cards, card)
	}

	if len(cards) != 5 {
		return Hand{}, fmt.Errorf("wrong card number %d", len(cards))
	}

	category, ranks := getCategoryAndRanks(cards)
	hand := Hand{
		input:    input,
		category: category,
		ranks:    ranks,
	}

	return hand, nil
}

func getCategoryAndRanks(cards Cards) (Category, []int) {
	category := HighCard

	ranks := make([]int, 0)
	for _, card := range cards {
		ranks = append(ranks, card.rank)
	}

	sort.Slice(ranks, func(a, b int) bool { return ranks[a] > ranks[b] })

	isStraight := true
	isFlush := true

	for i := 1; i < len(cards); i++ {
		if isFlush && cards[i-1].suit != cards[i].suit {
			isFlush = false
		}
		if isStraight && ranks[i-1] == 14 && ranks[i] == 5 {
			continue
		}
		if isStraight && ranks[i-1]-1 != ranks[i] {
			isStraight = false
		}
	}

	if isStraight {
		if ranks[0] == 14 && ranks[1] < 13 {
			ranks = append(ranks[1:], 1)
		}
		if isFlush {
			return StraightFlush, ranks
		}
		return Straight, ranks
	}

	if isFlush {
		category = Flush
	}

	rankCount := map[int]int{}
	for _, rank := range ranks {
		rankCount[rank]++
	}

	rankGroups := make([][2]int, 0)
	for rank, count := range rankCount {
		rankGroups = append(rankGroups, [2]int{rank, count})
	}

	sort.Slice(rankGroups, func(a, b int) bool {
		if rankGroups[a][1] == rankGroups[b][1] {
			return rankGroups[a][0] > rankGroups[b][0]
		}
		return rankGroups[a][1] > rankGroups[b][1]
	})

	ranks = make([]int, 0)

	for _, rankGroup := range rankGroups {
		rank, count := rankGroup[0], rankGroup[1]
		switch {
		case count == 4:
			category = FourOfKind
		case count == 3:
			category = ThreeOfKind
		case count == 2 && category == ThreeOfKind:
			category = FullHouse
		case count == 2 && category == OnePair:
			category = TwoPair
		case count == 2 && category == HighCard:
			category = OnePair
		}
		ranks = append(ranks, rank)
	}

	return category, ranks
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

	bestHand := hands[0]
	result := []string{bestHand.input}

	for _, hand := range hands[1:] {
		if bestHand.Equal(hand) {
			result = append(result, hand.input)
		}

		if bestHand.Less(hand) {
			result = []string{hand.input}
			bestHand = hand
		}
	}

	return result, nil
}
