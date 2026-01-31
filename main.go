package main

import (
	"fmt"

	"github.com/Cpt-Catnip/mtg-deck-util/pkg/deck"
)

func main() {
	d, err := deck.Import("./beeg_sqwirl.txt")
	if err != nil {
		fmt.Printf("error importing deck: %s", err)
	}
	fmt.Println(d.String())
	fmt.Println("~~~~~~~~~~~~~~~~~")
	fmt.Println(d.Size())
}
