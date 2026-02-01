package main

import (
	"fmt"

	"github.com/Cpt-Catnip/mtg-deck-util/pkg/deck"
)

func main() {
	d1, err := deck.Import("./beeg_sqwirl.txt")
	if err != nil {
		fmt.Printf("error importing deck: %s", err)
	}

	d2, err := deck.Import("./creative_squirrels_matter.txt")
	if err != nil {
		fmt.Printf("error importing deck: %s", err)
	}

	toAdd := d1.Diff(d2)
	toRemove := d2.Diff(d1)

	fmt.Printf("Cards to add\n============\n%s\n", toAdd)
	fmt.Printf("~~~~~~~~~~~~\n%d cards\n\n", toAdd.Size())
	fmt.Printf("Cards to remove\n===============\n%s\n", toRemove)
	fmt.Printf("~~~~~~~~~~~~\n%d cards\n", toRemove.Size())
}
