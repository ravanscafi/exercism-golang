package strand

var transcriptions = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	var rna string

	for _, n := range dna {
		rna += transcriptions[string(n)]
	}

	return rna
}
