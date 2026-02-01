package deck

import "testing"

func TestAdd(t *testing.T) {
	ts := []struct {
		name         string
		initialCards []string
		cardsToAdd   []string
		wantSize     int
	}{
		{
			name:         "should add one card",
			initialCards: []string{},
			cardsToAdd:   []string{"Llanowar Elves"},
			wantSize:     1,
		},
		{
			name:         "should add multiple cards",
			initialCards: []string{},
			cardsToAdd:   []string{"Llanowar Elves", "Chaos Warp", "Forest"},
			wantSize:     3,
		},
		{
			name:         "should add multiple of same card",
			initialCards: []string{},
			cardsToAdd:   []string{"Forest", "Forest", "Forest"},
			wantSize:     3,
		},
		{
			name:         "should add single additional card to existing set",
			initialCards: []string{"Forest", "Forest", "Forest"},
			cardsToAdd:   []string{"Forest"},
			wantSize:     4,
		},
		{
			name:         "should add multiple additional cards to existing set",
			initialCards: []string{"Forest", "Forest", "Forest"},
			cardsToAdd:   []string{"Forest", "Forest", "Forest"},
			wantSize:     6,
		},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			d := NewDeck(tc.initialCards...)
			for _, c := range tc.cardsToAdd {
				d.Add(c)
				assertDeckHas(t, d, c)
			}
			assertSizeEqual(t, d.size, tc.wantSize)
		})
	}

	// couldn't make this one fit neatly into the test cases
	t.Run("should update card count", func(t *testing.T) {
		d := NewDeck()
		cardToAdd := "Forest"
		count := 5
		for i := range count {
			d.Add(cardToAdd)
			assertCountEqual(t, d, cardToAdd, i+1)
		}
	})
}
