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

type Hand struct {
	cards    []Card
	input    string
	category Category
	score    int
	kickers  []int
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

	category, score, kickers := getCategory(hand.cards)
	hand.category = category
	hand.score = score
	hand.kickers = kickers

	return hand, nil
}

func getCategory(cards Cards) (Category, int, []int) {

	sort.Slice(cards, func(a, b int) bool { return cards[a].rank < cards[b].rank })

	firstCard := cards[0]
	totalValue := firstCard.rank
	isStraight := true
	isSameSuit := true
	streaks := [][]int{{firstCard.rank}}

	for i := 1; i < len(cards); i++ {
		totalValue += cards[i].rank
		if isStraight && cards[i-1].rank != cards[i].rank {
			isStraight = false
		}
		if isSameSuit && cards[i-1].suit != cards[i].suit {
			isSameSuit = false
		}
		if streaks[len(streaks)-1][0] == cards[i].rank {
			streaks[len(streaks)-1] = append(streaks[len(streaks)-1], cards[i].rank)
		} else {
			streaks = append(streaks, []int{cards[i].rank})
		}
	}

	if isStraight {
		if isSameSuit {
			return StraightFlush, totalValue, []int{}
		}
		return Straight, totalValue, []int{}
	}

	if isSameSuit {
		return Flush, totalValue, []int{}
	}

	pairs := 0
	score := 0
	category := HighCard
	kickers := make([]int, 0)

	for _, streak := range streaks {
		switch len(streak) {
		case 4:
			category = FourOfKind
			score += 4 * streak[0]
		case 3:
			category = ThreeOfKind
			score += 3 * streak[0]
		case 2:
			pairs++
			score += 2 * streak[0]
		default:
			kickers = append(kickers, streak[0])
		}
	}

	sort.Slice(kickers, func(a, b int) bool { return kickers[a] > kickers[b] })

	switch category {
	case FourOfKind:
		return category, score, kickers
	case ThreeOfKind:
		if pairs == 1 {
			return FullHouse, totalValue, kickers
		}
		return ThreeOfKind, totalValue, kickers
	default:
		if pairs == 2 {
			return TwoPair, score, kickers
		} else if pairs == 1 {

			return OnePair, score, kickers
		} else {
			return HighCard, 0, kickers
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
		aCategory, bCategory := hands[a].category, hands[b].category
		if aCategory == bCategory {
			return hands[a].score > hands[b].score
		}
		return aCategory > bCategory
	})

	highestCard := hands[0]

	result := make([]string, 0)

	for _, hand := range hands {
		if hand.category == highestCard.category && hand.score == highestCard.score {
			result = append(result, hand.input)
		}
	}

	return result, nil
}
