package cards

import (
	"math/rand"
)

const (
	Heart int = iota
	Club
	Diamond
	Spade
)

const (
	Ace int = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card uint8
type Deck struct {
	cards  [52]Card
	cursor int
}

var (
	suits    = []rune{'♥', '♣', '♦', '♠'}
	ranks    = []rune{'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K'}
	stringed = [52]string{}
)

func init() {
	for i := 0; i < 52; i++ {
		stringed[i] = string(ranks[Card(i).Rank()]) + string(suits[Card(i).Suit()])
	}
}

func NewDeck() Deck {
	ret := Deck{}
	for i := 0; i < 52; i++ {
		ret.cards[i] = Card(i)
	}
	return ret
}

func (c Card) Suit() int {
	return int(c) / 13
}

func (c Card) Rank() int {
	return int(c) % 13
}

func (c Card) String() string {
	return stringed[c]
}

func (d *Deck) IsExhausted() bool {
	return d.cursor == 51
}

func (d *Deck) Next() (Card, bool) {
	if d.cursor == 51 {
		return d.cards[d.cursor], true
	}
	d.cursor++
	return d.cards[d.cursor-1], false
}

func (d *Deck) Shuffle(seed int64) {
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < 52; i++ {
		o := r.Uint32() % 52
		d.cards[i], d.cards[o] = d.cards[o], d.cards[i]
	}
	d.cursor = 0
}
