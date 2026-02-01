package deck

import (
	"fmt"
	"maps"
	"strings"
)

func (d *Deck) Has(cardName string) bool {
	_, ok := d.cards[cardName]
	return ok
}

func (d *Deck) Get(cardName string) (Card, bool) {
	c, ok := d.cards[cardName]
	return c, ok
}

func (d *Deck) Set(cardName string, card Card) {
	d.cards[cardName] = card
}

func (d *Deck) Size() int {
	return d.size
}

func (d *Deck) String() string {
	//TODO: sort cards
	b := strings.Builder{}
	for _, c := range d.cards {
		fmt.Fprintf(&b, "%d %s\n", c.Count, c.Name)
	}
	s := b.String()
	return s[0 : len(s)-1]
}

func (d *Deck) Copy() *Deck {
	copy := NewDeck()
	copy.cards = maps.Collect(maps.All(d.cards))
	copy.size = d.size
	return copy
}
