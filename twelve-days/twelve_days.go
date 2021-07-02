package twelve

import "fmt"

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree.",
	2:  "two Turtle Doves, and ",
	3:  "three French Hens, ",
	4:  "four Calling Birds, ",
	5:  "five Gold Rings, ",
	6:  "six Geese-a-Laying, ",
	7:  "seven Swans-a-Swimming, ",
	8:  "eight Maids-a-Milking, ",
	9:  "nine Ladies Dancing, ",
	10: "ten Lords-a-Leaping, ",
	11: "eleven Pipers Piping, ",
	12: "twelve Drummers Drumming, ",
}

func Song() string {
	output := ""

	for i := 1; i <= 12; i++ {
		output += Verse(i)

		if i < 12 {
			output += "\n"
		}
	}

	return output
}

func Verse(day int) string {
	output := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[day])

	for i := day; i >= 1; i-- {
		output += gifts[i]
	}

	return output
}
