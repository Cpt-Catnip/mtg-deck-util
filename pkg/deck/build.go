package deck

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
	d.AddN(cardName, 1)
}

func (d *Deck) AddN(cardName string, count int) {
	if count < 1 {
		return
	}
	c, ok := d.cards[cardName]
	if !ok {
		d.cards[cardName] = Card{
			Name:  cardName,
			Count: count,
		}
	} else {
		c.Count += count
		d.Set(cardName, c)
	}
	d.size += count
}

func (d *Deck) Remove(cardName string) {
	d.RemoveN(cardName, 1)
}

func (d *Deck) RemoveN(cardName string, count int) {
	if count < 1 {
		return
	}
	c, ok := d.cards[cardName]
	if !ok {
		return
	}
	if count >= c.Count {
		delete(d.cards, cardName)
	} else {
		c.Count -= count
		d.Set(cardName, c)
	}
	d.size -= count
}

func (d *Deck) RemoveAll(cardName string) {
	delete(d.cards, cardName)
}
