package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type scoreCard struct {
	name           string
	mp, w, d, l, p int
}

//Tally tallies the scores of the given  teams and f them into a  table
func Tally(reader io.Reader, w io.Writer) error {
	scoreboard, err := parseAndTally(bufio.NewScanner(reader))
	if err != nil {
		return err
	}
	cards := make([]scoreCard, 0, len(scoreboard))
	for _, card := range scoreboard {
		cards = append(cards, card)
	}
	sort.SliceStable(cards, func(i, j int) bool {
		if cards[i].p == cards[j].p {
			return cards[i].name < cards[j].name
		}
		return cards[i].p > cards[j].p
	})
	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, card := range cards {
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", card.name, card.mp, card.w, card.d, card.l, card.p)
	}
	return nil
}

func parseAndTally(scanner *bufio.Scanner) (map[string]scoreCard, error) {
	scoreboard := make(map[string]scoreCard)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || text[0] == '#' {
			continue
		}
		splitText := strings.Split(text, ";")
		if len(splitText) != 3 {
			return nil, errors.New("incorrect format")
		}
		switch splitText[2] {
		case "win":
			name := splitText[0]
			card := scoreboard[name]
			card.name = name
			card.w++
			card.mp++
			card.p += 3
			scoreboard[name] = card
			name = splitText[1]
			card = scoreboard[name]
			card.name = name
			card.l++
			card.mp++
			scoreboard[name] = card
		case "loss":
			name := splitText[1]
			card := scoreboard[name]
			card.name = name
			card.w++
			card.mp++
			card.p += 3
			scoreboard[name] = card
			name = splitText[0]
			card = scoreboard[name]
			card.name = name
			card.l++
			card.mp++
			scoreboard[name] = card
		case "draw":
			name := splitText[0]
			card := scoreboard[name]
			card.name = name
			card.mp++
			card.p++
			card.d++
			scoreboard[name] = card
			name = splitText[1]
			card = scoreboard[name]
			card.name = name
			card.mp++
			card.p++
			card.d++
			scoreboard[name] = card
		default:
			return nil, errors.New("incorrect format")
		}
	}
	return scoreboard, nil
}
