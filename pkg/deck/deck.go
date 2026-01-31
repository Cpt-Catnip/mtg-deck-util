package deck

import (
	"fmt"
	"strings"
)

type Deck struct {
	cards map[string]Card
	size  int
}

type Card struct {
	Name  string
	Count int
}

func NewDeck(cs ...string) *Deck {
	d := Deck{
		cards: make(map[string]Card, 0),
		size:  0,
	}
	for _, c := range cs {
		d.Add(c)
	}
	return &d
}

func (d *Deck) Add(cardName string) {
	d.cards[cardName] = Card{
		Name:  cardName,
		Count: 1,
	}
	d.size = d.size + 1
}

func (d *Deck) AddN(cardName string, count int) {
	if count < 1 {
		return
	}
	d.cards[cardName] = Card{
		Name:  cardName,
		Count: count,
	}
	d.size += count
}

func (d *Deck) Remove(cardName string) {
	c, ok := d.cards[cardName]
	if !ok {
		return
	}
	if c.Count == 1 {
		delete(d.cards, cardName)
	} else {
		c.Count = c.Count - 1
	}
	d.size -= 1
}

func (d *Deck) Has(cardName string) bool {
	_, ok := d.cards[cardName]
	return ok
}

func (d *Deck) Size() int {
	return d.size
}

func (d *Deck) String() string {
	b := strings.Builder{}
	for _, c := range d.cards {
		fmt.Fprintf(&b, "%d %s\n", c.Count, c.Name)
	}
	s := b.String()
	return s[0 : len(s)-1]
}
