package deck

import "testing"

func TestAdd(t *testing.T) {
	ts := []struct {
		name  string
		cards []string
	}{
		{
			name:  "should add one card",
			cards: []string{"Llanowar Elves"},
		},
		{
			name:  "should add multiple cards",
			cards: []string{"Llanowar Elves", "Chaos Warp", "Forest"},
		},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDeck()
			for _, c := range tc.cards {
				d.Add(c)
				if !d.Has(c) {
					t.Errorf("could not find card %q in deck", c)
				}
			}

			if d.size != len(tc.cards) {
				t.Errorf("expected deck size to be %d but got %d", len(tc.cards), d.size)
			}
		})
	}
}
