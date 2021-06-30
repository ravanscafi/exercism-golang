package scale

var sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var majorScale = []string{"A", "B", "C", "D", "E", "F", "G"}
var useMajor = []string{"C", "a"}
var useSharps = []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
var useFlats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}


func Scale(tonic, interval string) []string {
	var after []string
	var scale = sharpScale

	if contains(useSharps, tonic) {
		scale = sharpScale
	}

	if contains(useFlats, tonic) {
		scale = flatScale
	}

	for i, note := range scale {
		if note == tonic {
			return append(scale[i:], after...)
		}
		after = append(after, note)
	}

	return nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}