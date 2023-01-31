package poker

type Suit int

const (
	Spade Suit = iota
	Club
	Diamond
	Heart
)

type Card struct {
	rank uint
	suit Suit
}

type Hand []Card

func BestHand(hands []string) ([]string, error) {
	panic("Please implement the BestHand function")
}
