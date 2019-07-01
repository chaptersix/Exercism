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
	name                                    string
	matchesPlayed, won, drawn, lost, points int
}

//Tally tallies the scores of the given  teams an fromates them into a  table
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
		if cards[i].points == cards[j].points {
			return cards[i].name < cards[j].name
		}
		return cards[i].points > cards[j].points
	})
	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, card := range cards {
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", card.name, card.matchesPlayed, card.won, card.drawn, card.lost, card.points)
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
			scoreboard[name] = scoreboard[name].win(name)
			name = splitText[1]
			scoreboard[name] = scoreboard[name].loss(name)

		case "loss":
			name := splitText[1]
			scoreboard[name] = scoreboard[name].win(name)
			name = splitText[0]
			scoreboard[name] = scoreboard[name].loss(name)

		case "draw":
			name := splitText[0]
			scoreboard[name] = scoreboard[name].draw(name)
			name = splitText[1]
			scoreboard[name] = scoreboard[name].draw(name)

		default:
			return nil, errors.New("incorrect format")
		}
	}
	return scoreboard, nil
}

func (card scoreCard) win(name string) scoreCard {
	card.name = name
	card.won++
	card.matchesPlayed++
	card.points += 3
	return card
}

func (card scoreCard) loss(name string) scoreCard {
	card.name = name
	card.lost++
	card.matchesPlayed++
	return card
}

func (card scoreCard) draw(name string) scoreCard {
	card.name = name
	card.matchesPlayed++
	card.points++
	card.drawn++
	return card
}
