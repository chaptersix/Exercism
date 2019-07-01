package twelve

import "fmt"

type dayOfChristmas struct {
	day, gift string
}

var lyrics = []dayOfChristmas{
	{"first", " a Partridge in a Pear Tree."},
	{"second", " two Turtle Doves, and"},
	{"third", " three French Hens,"},
	{"fourth", " four Calling Birds,"},
	{"fifth", " five Gold Rings,"},
	{"sixth", " six Geese-a-Laying,"},
	{"seventh", " seven Swans-a-Swimming,"},
	{"eighth", " eight Maids-a-Milking,"},
	{"ninth", " nine Ladies Dancing,"},
	{"tenth", " ten Lords-a-Leaping,"},
	{"eleventh", " eleven Pipers Piping,"},
	{"twelfth", " twelve Drummers Drumming,"},
}

var start = "On the %v day of Christmas my true love gave to me:%v%v"

//Song generates the song
func Song() string {
	sing := ""
	prev := ""
	for _, day := range lyrics {
		prev = day.gift + prev
		sing += fmt.Sprintf(start, day.day, prev, "\n")
	}
	return sing
}

// Verse returns a given line of the song
func Verse(line int) string {
	prev := ""
	for index, day := range lyrics {
		prev = day.gift + prev
		if index+1 == line {
			break
		}
	}
	return fmt.Sprintf(start, lyrics[line-1].day, prev, "")
}
