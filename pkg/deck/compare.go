package deck

import "maps"

func (d *Deck) Intersect(d2 *Deck) *Deck {
	intersection := NewDeck()
	for _, c := range d.cards {
		if d2.Has(c.Name) {
			count := min(c.Count, d2.cards[c.Name].Count)
			intersection.AddN(c.Name, count)
		}
	}
	return intersection
}

func (d *Deck) Union(d2 *Deck) *Deck {
	union := NewDeck()
	union.cards = maps.Collect(maps.All(d.cards))
	for _, c := range d2.cards {
		if !union.Has(c.Name) {
			union.AddN(c.Name, c.Count)
		} else {
			count1 := d.cards[c.Name].Count
			count2 := c.Count
			existingCard := union.cards[c.Name]
			existingCard.Count = max(count1, count2)
			union.cards[c.Name] = existingCard
		}
	}
	return union
}

func (d *Deck) Diff(d2 *Deck) *Deck {
	diff := NewDeck()
	diff.cards = maps.Collect(maps.All(d.cards))
	for _, c := range d2.cards {
		diff.RemoveN(c.Name, c.Count)
	}
	return diff
}
