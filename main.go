package main

import (
	"fmt"

	"github.com/Cpt-Catnip/mtg-deck-util/pkg/deck"
)

// TODO: add command line args
// TODO: figure out deck import but
func main() {
	d1, err := deck.Import("./heads_i_win.txt")
	if err != nil {
		fmt.Printf("error importing deck: %s", err)
	}

	d2, err := deck.Import("./mardu_surge.txt")
	if err != nil {
		fmt.Printf("error importing deck: %s", err)
	}

	// cards that are in common
	intersect := d1.Intersect(d2)
	fmt.Printf("Cards in both lists\n~~~~~~~~~~~~~~~~\n%s\n\n", intersect)

	// cards only in precon
	// justZurgo := d2.Diff(d1)
	// fmt.Printf("Cards only in Zurgo\n~~~~~~~~~~~~~~~~~~~~~~\n%s\n\n", justZurgo)
}
