package house


var phrases = []string{
	"house that Jack built.",
	"malt\nthat lay in the ",
	"rat\nthat ate the ",
	"cat\nthat killed the ",
	"dog\nthat worried the ",
	"cow with the crumpled horn\nthat tossed the ",
	"maiden all forlorn\nthat milked the ",
	"man all tattered and torn\nthat kissed the ",
	"priest all shaven and shorn\nthat married the ",
	"rooster that crowed in the morn\nthat woke the ",
	"farmer sowing his corn\nthat kept the ",
	"horse and the hound and the horn\nthat belonged to the ",
}

func Song() string {
	size := len(phrases)
	output := ""

	for i := 1; i < size; i++ {
		output += Verse(i) + "\n\n"
	}

	return output + Verse(size)
}

func Verse(n int) string {
	output := "This is the "

	for i := n - 1; i >= 0; i-- {
		output += phrases[i]
	}

	return output
}