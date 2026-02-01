package deck

import "testing"

func assertSizeEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected deck to have size %d but got %d", got, want)
	}
}

func assertDeckHas(t testing.TB, d *Deck, cardName string) {
	t.Helper()
	if !d.Has(cardName) {
		t.Errorf("expected deck to have %q but could not be found", cardName)
	}
}

func assertCountEqual(t testing.TB, d *Deck, cardName string, wantCount int) {
	t.Helper()
	c, ok := d.Get(cardName)
	if !ok {
		t.Errorf("card %q not in deck", cardName)
	}
	assertSizeEqual(t, c.Count, wantCount)
}
