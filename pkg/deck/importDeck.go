package deck

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: determine filetype
// TODO: determine format
func Import(filepath string) (*Deck, error) {
	d := NewDeck()

	// open deck list
	f, err := os.Open(filepath)
	if err != nil {
		return NewDeck(), fmt.Errorf("error opening deck list file: %w", err)
	}
	defer f.Close()

	// scan list line-by-line
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		elems := strings.SplitN(line, " ", 2)
		rawCount, cardName := elems[0], elems[1]
		count, err := strconv.Atoi(rawCount)
		if err != nil {
			return NewDeck(), fmt.Errorf("error reading count from %q", line)
		}
		d.AddN(cardName, count)
	}

	if err := s.Err(); err != nil {
		return NewDeck(), fmt.Errorf("error scanning deck list: %w", err)
	}

	return d, nil
}
